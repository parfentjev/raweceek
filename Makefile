-include .env
export

.PHONY: init fmt lint generate run

init:
	go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go get -tool github.com/sqlc-dev/sqlc/cmd/sqlc@latest

fmt:
	golangci-lint fmt

lint:
	golangci-lint run

generate:
	find internal/generated/ -iname '*.go' -delete
	go tool oapi-codegen --config spec/oapi-codegen-configuration.yaml spec/contract.yaml
	go tool sqlc -f spec/sqlc-configuration.yaml generate

run:
	go run cmd/server.go
