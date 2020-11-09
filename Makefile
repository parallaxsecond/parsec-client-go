# Copyright 2021 Contributors to the Parsec project.
# SPDX-License-Identifier: Apache-2.0

PROJECT_NAME := "parsec-client-go"
PKG := "github.com/parallaxsecond/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)


PROTOC_PREPROCESSED_FILES := $(shell find ./parsec-operations/protobuf -name '*.proto' -exec basename {} \; | awk '{print "wireinterface/go-protobuf/"$$1}')
PROTOC_OUTPUT_FILES=$(shell find parsec-operations/protobuf/ -name "*.proto" -exec basename {} .proto \; | awk '{print "wireinterface/operations/"$$1".pb.go"}')

.PHONY: all dep lint vet test test-coverage build  protoc protobuf_preprocess clean-protobuf clean clean-all
 
protobuf_preprocess: ${PROTOC_PREPROCESSED_FILES}

protoc: protobuf_preprocess ${PROTOC_OUTPUT_FILES} ## Generate protocol buffer go code

wireinterface/go-protobuf/%.proto: parsec-operations/protobuf/%.proto
	@mkdir -p wireinterface/go-protobuf
	@cp $< $@
	@$(eval PKG_NAME := $(shell basename $< .proto | sed s/_//g))
	@$(eval PKG_DEF := $(shell echo "option go_package = \\\"github.com/parallaxsecond/parsec-client-go/wireinterface/operations/$(PKG_NAME)\\\";"))
	@#echo gopkg $(PKG_DEF)
	@grep  "$(PKG_DEF)" $@ || echo "\n$(PKG_DEF)" >> $@


# Can't work out how to get path and filename into the match
# need to have operations/option/option.pb.go maping to interface/go-protobuf/option.proto
# But works quickly and not needed often
wireinterface/operations/%.pb.go: wireinterface/go-protobuf/%.proto
	@protoc -I=wireinterface/go-protobuf --go_out=../../../ $< > /dev/null

clean-all: clean clean-protobuf
clean:
	@go clean ./...
	@rm -f $(PROJECT_NAME)/buildmk	

clean-protobuf:
	@find wireinterface/operations/ -name "*.pb.go" -exec rm {} \;
	@rm -Rf wireinterface/go-protobuf


all: protoc build ## Generate protocol buffer code and compile

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

sec: ## Run gosec for security analysis
	@gosec -quiet ./...

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

ci-test-all: ## Run Continuous Integration tests for all providers
	@./e2etest/ci-all.sh

build: dep ## Build the binary file
	@go build -i -o ./... $(PKG)
 

	
 
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'