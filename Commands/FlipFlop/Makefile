GOPATH:=$(shell go env GOPATH)
PROTO:= "../../Commands/FlipFlop/proto"

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: proto
proto:
	@protoc --proto_path=../../ \
		--go_out=plugins=grpc:${PROTO}/ \
		../../Shared/proto/block.proto \
		${PROTO}/commands.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o frontend *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build --build-arg DIR=Commands/FlipFlop -t FlipFlop:latest -f ../../Commands/FlipFlop/dockerfile ../../
