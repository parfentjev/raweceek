-include .env
export

.PHONY: init fmt lint generate tsc test run

init:
	@command -v go >/dev/null 2>&1 || { echo "error: 'go' is not installed"; exit 1; }
	go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go get -tool github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@command -v tsc >/dev/null 2>&1 || echo "warning: 'typescript' is not installed"

fmt:
	golangci-lint fmt

lint:
	golangci-lint run

generate:
	find internal/generated/ -iname '*.go' -delete
	go tool oapi-codegen --config spec/oapi-codegen-configuration.yaml spec/contract.yaml
	go tool sqlc -f spec/sqlc-configuration.yaml generate

tsc:
	tsc --project tsconfig.json

test:
	go test -v ./...

run:
	go run cmd/server.go
