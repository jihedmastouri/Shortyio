.PHONY: init
init:
	@docker pull kjconroy/sqlc

.PHONY: proto
proto:
	@protoc --proto_path=./proto/ \
		--go_out=plugins=grpc:. \
		./proto/block.proto \
		./proto/queries.proto \
		./proto/commands.proto

.PHONY: codegen
codegen:
	@docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

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
