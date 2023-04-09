.PHONY: init
init:
	@docker pull kjconroy/sqlc

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
