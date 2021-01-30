.PHONY: build test

build:
	go build -o nntp ./cmd/main.go

test:
	go test ./...
