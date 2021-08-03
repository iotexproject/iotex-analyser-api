
########################################################################################################################
# Copyright (c) 2020 IoTeX
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

.PHONY: build proto all run

all : build

proto:
	#protoc -I ./proto --go_out ./  --go-grpc_out ./ --grpc-gateway_out ./ proto/*.proto
	protoc -I ./proto --go_out ./ --go-grpc_out ./ --grpc-gateway_out ./ --graphql_out ./ proto/api_actions.proto
	rm -f api/api_actions.graphql.go && mv api/api.graphql.go api/api_actions.graphql.go

clean:
	rm -f iotex-analyser-api
	
build:
	$(GOBUILD) -v .

run: build
