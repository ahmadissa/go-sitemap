## Setup
setup:
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/mattn/go-colorable

## test
test:
	go test -v -cover ./...

## lint
lint:
	golangci-lint run ./...

## benchmark
benchmark:
	go test -bench . -benchmem

## Show help
help:
	@make2help $(MAKEFILE_LIST)
