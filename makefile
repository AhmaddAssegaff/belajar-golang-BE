.PHONY: dev test build run tidy swag

PKG ?= ./...
ARGS ?=

dev:
	go run ./cmd

test:
	go test -v $(PKG) $(ARGS)

build:
	go build -o app ./cmd

run: build
	./app

tidy:
	go mod tidy

swag:
	swag init
