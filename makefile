.PHONY: dev test build run tidy swag migrate-up migrate-down migrate-create migrate-seed

include .env
export

PKG ?= ./...
ARGS ?=

dev:
	go run ./cmd

unit-test:
	go test -v $(PKG) $(ARGS)

e2e-test:
	go test ./test/e2e -v $(if $(FILE),./test/e2e/$(FILE)) $(if $(RUN),-run $(RUN))

build:
	go build -o app ./cmd

run: build
	./app

tidy:
	go mod tidy

swag:
	swag init -g cmd/main.go -o docs

sqlc-gen:
	sqlc generate

migrate-create:
	migrate create -ext sql -dir ./migrations/ -seq $(name)

migrate-up:
	migrate -path ./migrations/ -database "$(DB_URL)" up

migrate-up-one:
	migrate -path ./migrations/ -database "$(DB_URL)" up 1

migrate-down:
	migrate -path ./migrations/ -database "$(DB_URL)" down 1

migrate-down-all:
	migrate -path ./migrations/ -database "$(DB_URL)" down

migrate-force:
	migrate -path ./migrations/ -database "$(DB_URL)" force $(version)

migrate-version:
	migrate -path ./migrations/ -database "$(DB_URL)" version
