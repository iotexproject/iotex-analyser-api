# iotex-kit 迁移新增接口 — 说明

> 本文档说明为「iotex-kit 从直连 analyzer Postgres 迁移到调用本 API」而在本仓库
> 新增的接口、如何本地运行/测试、以及线上部署是否需要改动。

## 1. 新增了什么

为承接 iotex-kit 原本直连 analyzer DB 的查询，本仓库新增 3 组共 19 个 RPC（全部 HTTP POST +
gRPC-gateway，路径 `POST /<package>.<Service>.<Method>`）。

### 新增文件
- `proto/api_iotexscan.proto` — `IotexscanService`（etherscan 兼容端点，11 RPC）
- `apiservice/iotexscan_service.go` — `IotexscanService` 实现
- `apiservice/delegate_migration_service.go` — `DelegateService` 追加的 7 个方法
- `apiservice/staking_migration_service.go` — `StakingService` 追加的 `GetStakingHistory`
- `apiservice/migration_integration_test.go` — 上述接口的集成测试（对拍真实 DB）
- 生成物：`api/api_iotexscan*.go`（新），`api/api_delegate*.go` / `api/api_staking*.go`（重新生成）

### 改动文件
- `proto/api_delegate.proto` / `proto/api_staking.proto` — 追加迁移 RPC
- `apiservice/apiserver.go` — 注册 `IotexscanService`（gRPC + gateway handler）
- `Makefile` — `proto` target 增加 `api_iotexscan.proto` 一行

### RPC 清单
| Service | 新增 RPC |
|---------|----------|
| IotexscanService | GetTxListByAddress, GetTokenTxByAddress, GetTokenNftTxByAddress, GetToken1155TxByAddress, GetTxListInternal, GetContractLogs, GetGasOracle, GetDailyNewAddresses, GetContractCreationBatch, GetBlockNumberByTime, GetActionStatusByHash |
| DelegateService | GetDelegateHeight, GetProductivityHistory, GetProbationHistory, GetDelegateRewards, GetDelegateRewardsHistory, GetReceivedVotesByAddress, GetDelegatesStatistics |
| StakingService | GetStakingHistory |

## 2. 重新生成 proto（改了 .proto 时）

需要 protoc + 4 个插件（本地一次性安装）：

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql@latest   # delegate/staking 有 graphql 注解，必须装

export PATH="$PATH:$(go env GOPATH)/bin"
make proto
```

注意：`api_iotexscan.proto` 只用 `google.api.http`（无 graphql 注解），Makefile 里它那行不带
`--graphql_out`；delegate/staking 仍需 graphql 插件生成已有的 graphql handler。

## 3. 本地构建 / 运行

### 构建
```bash
go build ./...          # 或 make
go vet ./apiservice/
```

### 本地调试运行（连远程 analyzer DB）
提供了独立的调试用文件，不影响生产 `Dockerfile`：
- `Dockerfile.dev` — 单阶段调试镜像（默认走 goproxy.cn 便于国内构建）
- `docker-compose.dev.yml` — 挂载宿主 Go module 缓存 + 源码，容器内编译运行（离线、改代码重启即重编译）
- `.env.docker.example` — 复制为 `.env.docker`（已 gitignore）填远程 DB 凭据

```bash
cp .env.docker.example .env.docker      # 填 DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME
docker compose -f docker-compose.dev.yml up -d analyser
curl http://localhost:8889/healthz      # 应返回 ok
```
HTTP 8889（kit 调用）、gRPC 8888。DB 若在本机，`.env.docker` 里 `DB_HOST=host.docker.internal`。
只读 standby 建议 `DB_SKIP_AUTO_MIGRATE=true`（AutoMigrate 会发 DDL，standby 拒写）。

## 4. 运行测试

### 单元/纯函数测试（不连 DB，随时可跑）
```bash
go test ./apiservice/ -run 'TestParseDateRange|Test.*_Unit'
```

### 集成测试（对拍真实 DB，默认跳过）
`apiservice/migration_integration_test.go` 与 `new_endpoints_integration_test.go` 采用同一模式：
仅当 `ITEST_DB_HOST` 设置时才运行，否则 `t.Skip`，绝不在 CI 误连。

```bash
ITEST_DB_HOST=<host> ITEST_DB_PORT=5432 \
ITEST_DB_USER=<user> ITEST_DB_PASSWORD=<pw> ITEST_DB_NAME=mainnet \
  go test ./apiservice -run TestMigration_Integration -count=1 -v
```
覆盖：GetBlockNumberByTime(before/after)、GetDailyNewAddresses、GetActionStatusByHash、
GetContractCreationBatch（含 `IN(?)` 回归）、GetDelegatesStatistics、GetStakingHistory、
GetReceivedVotesByAddress、GetProductivityHistory 校验。

## 5. 线上部署 — 是否需要改环境变量？

**不需要新增环境变量。** 新接口复用现有配置：

- **DB**：沿用现有 `DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME`（同一个 analyzer 库）。
- **鉴权**：新接口走调用方的服务 key JWT，**不加入** `auth/middleware.go` 的 `whitelistAPI`
  （那是完全免鉴权列表，只给公开端点）。
- **限流豁免**：`auth/ratelimit.go` 的 `whitelistID` 按 **JWT 身份**豁免（非按路由），iotex-kit 用的
  服务 key 已在其中，所以新 RPC 自动绕过 5 req/min，**无需为每个新 RPC 改 auth**。

上线动作只需：合并代码 → 常规构建部署（生产 `Dockerfile` 未变，仍 `make` 构建）。

## 6. 已知遗留 / 注意

- **`delegate_rewards` 表**：`GetDelegateRewards` 的 SQL 1:1 沿用 kit 原代码查 `delegate_rewards`。
  该表此前在主网库缺失（历史数据迁移遗留），现已恢复（`mainnet` 库 123 行），该接口已正常返回数据。
  ✅ **已解决**（2026-07-02 验证：真实 candidate 返回完整奖励字段，缺失 candidate 返回空且不报错）。
- **表名对齐**：`block_action` 实为 `block_action_partition`；ERC721 转账实为 `erc721_transfers_v2_2_3`；
  ERC1155 为 `erc1155_transfer_singles_v2_2_2`。相关 SQL 已按真实表名写。
- **慢查询**：`erc20_transfers`（2 亿+行）的 token 转账查询用两腿 UNION（各走 sender/recipient 索引），
  避免 `(sender OR recipient) ORDER BY` 全表反扫超时。
- **批量参数**：GORM `Raw` 传 Go slice 用 `IN (?)`（展开 `IN ($1,...)`），不能用 `ANY(?)`
  （不渲染成 PG array literal，报 22P02）。

## 7. 真实调用验证结果

对拍真实主网库（`mainnet`）逐个实调，**19 个新增 RPC + 复用的 3 个，22/22 全通过**。
（原唯一失败项 `GetDelegateRewards` 因 `delegate_rewards` 表缺失，该表恢复后已通过。）

| RPC | 结果 | 备注 |
|-----|------|------|
| IotexscanService.GetTxListByAddress | ✅ | 真实数据，etherscan 字段完整 |
| IotexscanService.GetTxListInternal | ✅ | 真实内部交易 |
| IotexscanService.GetTokenTxByAddress (ERC20) | ✅ | 修慢查询（UNION）后快速返回 |
| IotexscanService.GetTokenNftTxByAddress (ERC721) | ✅ | 修表名 `erc721_transfers_v2_2_3` 后返回 |
| IotexscanService.GetToken1155TxByAddress | ✅ | 空结果但 SQL 正常 |
| IotexscanService.GetContractLogs | ✅ | 真实日志（topic/data/hash） |
| IotexscanService.GetGasOracle | ✅ | 逻辑正常（该库 store 无此 key） |
| IotexscanService.GetDailyNewAddresses | ✅ | 入参日期区间生效 |
| IotexscanService.GetContractCreationBatch | ✅ | 修 `ANY(?)`→`IN(?)` 后返回创建者/创建 tx |
| IotexscanService.GetBlockNumberByTime | ✅ | before/after 双向 |
| IotexscanService.GetActionStatusByHash | ✅ | action + bucket |
| DelegateService.GetDelegateHeight | ✅ | 真实高度 |
| DelegateService.GetProductivityHistory | ✅ | SQL 正常 |
| DelegateService.GetProbationHistory | ✅ | 真实 probation 记录 |
| DelegateService.GetDelegateRewardsHistory | ✅ | SQL 正常 |
| DelegateService.GetReceivedVotesByAddress | ✅ | 真实 staker/amount |
| DelegateService.GetDelegatesStatistics | ✅ | 123 delegates + 总质押量 |
| DelegateService.GetDelegateRewards | ✅ | `delegate_rewards` 表已恢复（123 行），真实奖励数据 |
| StakingService.GetStakingHistory | ✅ | 真实 bucket |
| ChainService.GetBlockMeta（复用） | ✅ | block.getblockreward |
| AccountService.GetContractByteCode（复用） | ✅ | 真实字节码 |
| StakingService.GetBucketByActionHash（复用） | ✅ | getstatus 的 bucket |

kit 侧同时验证：mock 回归测试 `bun test` 21/21 通过；端到端经 kit `/api` 打通（txlist / getblockreward /
getblocknobytime / dailynewaddress / getstatus / gettxreceiptstatus / getcontractcreation 等）。

对外 API 契约**未变更**：所有接口的入参（zod schema）与响应字段名（含 `txreceipt_status` /
`tokenID` / bucket 的 snake_case 等特殊命名）均保持与迁移前一致，仅数据来源从 analyzer 直连改为本 API。
