# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
make              # Build the binary
make run          # Build and run the server
make proto        # Regenerate protobuf/gRPC/GraphQL code from .proto files
make docker       # Build Docker image

go test ./...           # Run all tests
go test -v ./path/...   # Run tests in a specific package (verbose)
go vet ./...            # Static analysis
```

**Always run `make proto` after modifying any `.proto` file.** Generated code lives in `/api/` and must not be edited by hand.

## Architecture

This is a multi-protocol (gRPC + HTTP REST + GraphQL) IoTeX blockchain analytics API. It queries a PostgreSQL/MySQL/SQLite database populated by an external indexer.

**Two servers start concurrently:**
- Port 8888: native gRPC
- Port 8889: HTTP/REST (gRPC-gateway) + GraphQL playground

**Request flow:** HTTP request → gRPC-gateway proxy → gRPC service handler → GORM query → DB

### Layering

| Layer | Location | Role |
|-------|----------|------|
| Protocol definitions | `/proto/` | Source of truth for all API shapes |
| Generated bindings | `/api/` | gRPC stubs, gateway, GraphQL — never edit |
| Service implementations | `/apiservice/` | Implement generated gRPC interfaces |
| Business logic | `/common/` | Reusable query helpers (accounts, actions, votings, rewards) |
| Database | `/db/` | GORM connection manager |
| Models | `/model/` | GORM struct definitions |
| Auth | `/auth/` | JWT middleware, whitelist, rate limiting |
| Config | `/config/` | YAML + env config loading |

### Adding a New Endpoint

1. Define the RPC in the relevant `.proto` file (add HTTP + GraphQL annotations)
2. `make proto` to regenerate `/api/`
3. Implement the method in the corresponding `/apiservice/*_service.go`
4. Add business logic helpers in `/common/` if needed
5. Register on the whitelist in `/auth/` if unauthenticated access is required

### Key Conventions

- **Pagination**: all list endpoints use `common.Pagination` — see `/common/common.go` for helpers
- **Address validation**: use `common.Address()` from `/common/common.go`; IoTeX addresses use the `io1...` format
- **Error handling**: return gRPC status errors; avoid panics in service handlers
- **Config precedence**: env vars override YAML, which overrides hardcoded defaults (`GRPC_API_PORT`, `HTTP_API_PORT`, `DB_*`, `CHAIN_GRPC_ENDPOINT`)
- **Multi-DB support**: queries must be compatible with PostgreSQL, MySQL, and SQLite3 via GORM

### Key Dependencies

- `google.golang.org/grpc` + `grpc-ecosystem/grpc-gateway/v2` — transport
- `ysugimoto/grpc-graphql-gateway` — GraphQL generation
- `gorm.io/gorm` — ORM (drivers: postgres, mysql, sqlite)
- `iotexproject/iotex-core/v2` + `iotex-antenna-go/v2` — IoTeX types and SDK
