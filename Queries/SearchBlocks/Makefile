GOPATH:=$(shell go env GOPATH)
PATH:="../../Queries/SearchBlocks/proto"

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: proto
proto:
	@protoc --proto_path=../../ \
		--go_out=plugins=grpc:{PATH}/ \
		../../Shared/proto/block.proto \
		{PATH}/queries.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o service *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build --build-arg DIR=Queries/SearchBlocks -t SearchBlocks:latest -f ../../Queries/SearchBlocks/dockerfile ../../
