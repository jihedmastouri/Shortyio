GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: proto
proto:
	@protoc --proto_path=../../ \
		--go_out=plugins=grpc:../../Commands/FlipFlop/proto/ \
		../../Shared/proto/block.proto \
		../../Commands/FlipFlop/proto/commands.proto

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
