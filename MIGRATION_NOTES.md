# iotex-kit Migration — New Endpoints Notes

> This document describes the endpoints added to this repo to support
> "migrating iotex-kit from a direct analyzer-Postgres connection to calling this API",
> how to run/test locally, and whether production deployment needs any changes.

## 1. What was added

To take over the queries iotex-kit used to run directly against the analyzer DB, this repo adds
3 groups totaling 19 RPCs (all HTTP POST + gRPC-gateway, path `POST /<package>.<Service>.<Method>`).

### New files
- `proto/api_iotexscan.proto` — `IotexscanService` (etherscan-compatible endpoints, 11 RPCs)
- `apiservice/iotexscan_service.go` — `IotexscanService` implementation
- `apiservice/delegate_migration_service.go` — 7 methods appended to `DelegateService`
- `apiservice/staking_migration_service.go` — `GetStakingHistory` appended to `StakingService`
- `apiservice/migration_integration_test.go` — integration tests for the above (against a real DB)
- Generated: `api/api_iotexscan*.go` (new), `api/api_delegate*.go` / `api/api_staking*.go` (regenerated)

### Modified files
- `proto/api_delegate.proto` / `proto/api_staking.proto` — appended migration RPCs
- `apiservice/apiserver.go` — register `IotexscanService` (gRPC + gateway handler)
- `Makefile` — added an `api_iotexscan.proto` line to the `proto` target

### RPC list
| Service | New RPCs |
|---------|----------|
| IotexscanService | GetTxListByAddress, GetTokenTxByAddress, GetTokenNftTxByAddress, GetToken1155TxByAddress, GetTxListInternal, GetContractLogs, GetGasOracle, GetDailyNewAddresses, GetContractCreationBatch, GetBlockNumberByTime, GetActionStatusByHash |
| DelegateService | GetDelegateHeight, GetProductivityHistory, GetProbationHistory, GetDelegateRewards, GetDelegateRewardsHistory, GetReceivedVotesByAddress, GetDelegatesStatistics |
| StakingService | GetStakingHistory |

## 2. Regenerating proto (when a .proto changes)

Requires protoc + 4 plugins (one-time local install):

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql@latest   # delegate/staking have graphql annotations, required

export PATH="$PATH:$(go env GOPATH)/bin"
make proto
```

Note: `api_iotexscan.proto` uses only `google.api.http` (no graphql annotations), so its Makefile line
omits `--graphql_out`; delegate/staking still need the graphql plugin to generate their existing graphql handlers.

## 3. Local build / run

### Build
```bash
go build ./...          # or make
go vet ./apiservice/
```

### Local debug run (against a remote analyzer DB)
Standalone debug files are provided; they do NOT affect the production `Dockerfile`:
- `Dockerfile.dev` — single-stage debug image (defaults to goproxy.cn for easier builds in China)
- `docker-compose.dev.yml` — mounts the host Go module cache + source and compiles/runs inside the container (offline; recompiles on restart after code changes)
- `.env.docker.example` — copy to `.env.docker` (gitignored) and fill in the remote DB credentials

```bash
cp .env.docker.example .env.docker      # fill DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME
docker compose -f docker-compose.dev.yml up -d analyser
curl http://localhost:8889/healthz      # should return ok
```
HTTP 8889 (called by kit), gRPC 8888. If the DB is on the host machine, set `DB_HOST=host.docker.internal`
in `.env.docker`. For a read-only standby, `DB_SKIP_AUTO_MIGRATE=true` is recommended (AutoMigrate issues
DDL, which a standby rejects).

## 4. Running tests

### Unit / pure-function tests (no DB, always runnable)
```bash
go test ./apiservice/ -run 'TestParseDateRange|Test.*_Unit'
```

### Integration tests (against a real DB, skipped by default)
`apiservice/migration_integration_test.go` and `new_endpoints_integration_test.go` follow the same pattern:
they run only when `ITEST_DB_HOST` is set, otherwise `t.Skip`, so they never accidentally connect in CI.

```bash
ITEST_DB_HOST=<host> ITEST_DB_PORT=5432 \
ITEST_DB_USER=<user> ITEST_DB_PASSWORD=<pw> ITEST_DB_NAME=mainnet \
  go test ./apiservice -run TestMigration_Integration -count=1 -v
```
Covers: GetBlockNumberByTime (before/after), GetDailyNewAddresses, GetActionStatusByHash,
GetContractCreationBatch (incl. the `IN(?)` regression), GetDelegatesStatistics, GetStakingHistory,
GetReceivedVotesByAddress, and GetProductivityHistory validation.

## 5. Production deployment — any env var changes needed?

**No new environment variables required.** The new endpoints reuse existing config:

- **DB**: reuses the existing `DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME` (same analyzer DB).
- **Auth**: new endpoints go through the caller's service-key JWT; they are **not added** to `whitelistAPI`
  in `auth/middleware.go` (that is the fully-unauthenticated list, only for public endpoints).
- **Rate-limit exemption**: `whitelistID` in `auth/ratelimit.go` exempts by **JWT identity** (not by route);
  the service key iotex-kit uses is already in it, so new RPCs automatically bypass the 5 req/min limit —
  **no auth change is needed per new RPC**.

Rollout is just: merge code → normal build & deploy (the production `Dockerfile` is unchanged, still built with `make`).

## 6. Known caveats / notes

- **`delegate_rewards` table**: `GetDelegateRewards`'s SQL is a 1:1 port of kit's original code querying
  `delegate_rewards`. The table was previously missing from the mainnet DB (a historical data-migration
  leftover), but has now been restored (`mainnet` DB, 123 rows), and the endpoint returns data normally.
  ✅ **Resolved** (verified 2026-07-02: a real candidate returns full reward fields; a missing candidate
  returns empty without error).
- **Table-name alignment**: `block_action` is actually `block_action_partition`; ERC721 transfers are
  `erc721_transfers_v2_2_3`; ERC1155 is `erc1155_transfer_singles_v2_2_2`. The relevant SQL uses the real table names.
- **Slow queries**: the token-transfer query on `erc20_transfers` (200M+ rows) uses a two-leg UNION
  (each leg using the sender/recipient index) to avoid a full-table reverse scan timeout from
  `(sender OR recipient) ORDER BY`.
- **Batch params**: passing a Go slice to GORM `Raw` uses `IN (?)` (expands to `IN ($1,...)`); do NOT use
  `ANY(?)` (it does not render as a PG array literal and errors with 22P02).

## 7. Real-call verification results

Called against the real mainnet DB (`mainnet`) one by one: **19 new RPCs + 3 reused, 22/22 all passing.**
(The originally-failing `GetDelegateRewards` failed because the `delegate_rewards` table was missing;
it now passes after the table was restored.)

| RPC | Result | Notes |
|-----|--------|-------|
| IotexscanService.GetTxListByAddress | ✅ | Real data, full etherscan fields |
| IotexscanService.GetTxListInternal | ✅ | Real internal transactions |
| IotexscanService.GetTokenTxByAddress (ERC20) | ✅ | Fast after the slow-query fix (UNION) |
| IotexscanService.GetTokenNftTxByAddress (ERC721) | ✅ | Returns after the table-name fix `erc721_transfers_v2_2_3` |
| IotexscanService.GetToken1155TxByAddress | ✅ | Empty result but SQL is correct |
| IotexscanService.GetContractLogs | ✅ | Real logs (topic/data/hash) |
| IotexscanService.GetGasOracle | ✅ | Logic correct (this DB's store has no such key) |
| IotexscanService.GetDailyNewAddresses | ✅ | Date-range input takes effect |
| IotexscanService.GetContractCreationBatch | ✅ | Returns creator/creation tx after the `ANY(?)`→`IN(?)` fix |
| IotexscanService.GetBlockNumberByTime | ✅ | before/after both directions |
| IotexscanService.GetActionStatusByHash | ✅ | action + bucket |
| DelegateService.GetDelegateHeight | ✅ | Real height |
| DelegateService.GetProductivityHistory | ✅ | SQL correct |
| DelegateService.GetProbationHistory | ✅ | Real probation records |
| DelegateService.GetDelegateRewardsHistory | ✅ | SQL correct |
| DelegateService.GetReceivedVotesByAddress | ✅ | Real staker/amount |
| DelegateService.GetDelegatesStatistics | ✅ | 123 delegates + total stake |
| DelegateService.GetDelegateRewards | ✅ | `delegate_rewards` table restored (123 rows), real reward data |
| StakingService.GetStakingHistory | ✅ | Real bucket |
| ChainService.GetBlockMeta (reused) | ✅ | block.getblockreward |
| AccountService.GetContractByteCode (reused) | ✅ | Real bytecode |
| StakingService.GetBucketByActionHash (reused) | ✅ | The bucket for getstatus |

Verified on the kit side too: mock regression tests `bun test` 21/21 pass; end-to-end through kit `/api`
(txlist / getblockreward / getblocknobytime / dailynewaddress / getstatus / gettxreceiptstatus /
getcontractcreation, etc.).

The external API contract is **unchanged**: every endpoint's inputs (zod schema) and response field names
(including special casing like the snake_case `txreceipt_status` / `tokenID` / bucket fields) remain
identical to before the migration; only the data source changed from a direct analyzer connection to this API.
