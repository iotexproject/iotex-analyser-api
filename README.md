[![Build Status](https://github.com/iotexproject/iotex-analyser-api.svg?branch=main)](https://travis-ci.org/iotexproject/iotex-analyser-api)

# Overview
API for iotex-analyser

## Build from Code

Install Google protocol buffers compiler [protoc](https://github.com/protocolbuffers/protobuf) 

***install required protoc plugins***: 
```
go install github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql/...
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.4
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
```

Compiling protocol buffers and generate Go code, Only when you modified proto file.
```
make proto
```

Build Server
```
make
```

## Docker Quick Start

API service depends on [iotex-analyser](https://github.com/iotexproject/iotex-analyser). You need to use docker start `iotex-analyser` before that.

```
docker run -p 8888:8888 -p 8889:8889 -e "GRPC_API_PORT=8888" -e "HTTP_API_PORT=8889" -e "CHAIN_GRPC_ENDPOINT=api.iotex.one:443" -e "DB_DRIVER=postgres" -e "DB_HOST=x.x.x.x" -e "DB_PORT=5432" -e "DB_USER=user" -e "DB_PASSWORD=password" -e "DB_NAME=dbname" iotexproject/iotex-analyser-api
```
* Note: Please change your database config in command 




API supports gRPC,RESTful HTTP and GraphQL

```sh
curl -g "http://localhost:8889/graphql" -d '
{
  GetActionsByAddress(address: "io14u5d66rt465ykm7t2847qllj0reml27q30kr75") {
    count
    results{
      actHash
      amount
    }
  }
}'

curl -g "http://localhost:8889/api.ActionsService.GetActionsByAddress" -d '
{
  "address": "io14u5d66rt465ykm7t2847qllj0reml27q30kr75"
}'

grpcurl -plaintext -d '{"address": "io14u5d66rt465ykm7t2847qllj0reml27q30kr75"}' 127.0.0.1:8888 api.ActionsService.GetActionsByAddress
```

GraphQL Playground Support

http://localhost:8889/graphql

## License
This project is licensed under the [Apache License 2.0](LICENSE).
