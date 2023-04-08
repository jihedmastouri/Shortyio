GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --go_out=:. proto/hipstershop.proto

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
	@docker build -t frontend:latest .

.PHONY: generate
generate:
	@docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate
