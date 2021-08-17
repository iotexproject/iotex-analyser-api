[![Build Status](https://github.com/iotexproject/iotex-analyser-api.svg?branch=main)](https://travis-ci.org/iotexproject/iotex-analyser-api)

# Overview
API for iotex-analyser

## Build from Code

Install Google protocol buffers compiler [protoc](https://github.com/protocolbuffers/protobuf) 

***install required protoc plugins***: 
```
go get github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql/...
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.4
```

Compiling protocol buffers and generate Go code, Only when you modified proto file.
```
make proto
```

Build Server
```
make
```

API supports GRPC/HTTP/GraphQL

```sh
curl -g "http://localhost:7778/graphql" -d '
{
  GetActionsByAddress(address: "io14u5d66rt465ykm7t2847qllj0reml27q30kr75") {
    count
    results{
      actHash
      amount
    }
  }
}'

curl -g "http://localhost:7778/api.ActionsService.GetActionsByAddress" -d '
{
  "address": "io14u5d66rt465ykm7t2847qllj0reml27q30kr75"
}'

grpcurl -plaintext -d '{"address": "io14u5d66rt465ykm7t2847qllj0reml27q30kr75"}' 127.0.0.1:7777 api.ActionsService.GetActionsByAddress
```

## License
This project is licensed under the [Apache License 2.0](LICENSE).
