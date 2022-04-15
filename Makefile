
########################################################################################################################
# Copyright (c) 2020 IoTeX
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

NAME=iotex/iotex-analyser-api
# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
CWD=$(abspath $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

.PHONY: build doc docs proto all run

all : build

proto:
	protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/include/pagination.proto
	rm -rf api/pagination
	mv github.com/iotexproject/iotex-analyser-api/api/pagination api/
	rm -rf github.com/
	protoc -I ./proto --go_out ./ --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_actions.proto
	rm -f api/api_actions.graphql.go && mv api/api.graphql.go api/api_actions.graphql.go
	protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_staking.proto
	rm -f api/api_staking.graphql.go && mv api/api.graphql.go api/api_staking.graphql.go
	protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_account.proto
	rm -f api/api_account.graphql.go && mv api/api.graphql.go api/api_account.graphql.go
	protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_delegate.proto
	rm -f api/api_delegate.graphql.go && mv api/api.graphql.go api/api_delegate.graphql.go
	protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_chain.proto
	rm -f api/api_chain.graphql.go && mv api/api.graphql.go api/api_chain.graphql.go
	protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_action.proto
	rm -f api/api_action.graphql.go && mv api/api.graphql.go api/api_action.graphql.go
clean:
	rm -f iotex-analyser-api
	
doc:
	protoc -I  ./proto --doc_out=./doc  --doc_opt=html,docs.html proto/*.proto proto/include/pagination.proto
	protoc -I  ./proto --doc_out=./doc  --doc_opt=markdown,readme.md proto/*.proto proto/include/pagination.proto
	protoc -I  ./proto --doc_out=./doc  --doc_opt=docbook,docs.xml proto/*.proto proto/include/pagination.proto
	protoc -I  ./proto --doc_out=./doc  --doc_opt=json,docs.json proto/*.proto proto/include/pagination.proto

docs:
	# docuowl --input docs --output docs-html
	docker run --rm --name slate -v $(CWD)/docs-html:/srv/slate/build -v $(CWD)/docs:/srv/slate/source slatedocs/slate build
build:
	$(GOBUILD) -v .

run: build

docker:
	docker build --progress=plain -t ${NAME}:latest  .
