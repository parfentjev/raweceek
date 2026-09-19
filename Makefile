-include .env
export

.PHONY: fmt lint tsc run

fmt:
	cargo fmt

lint:
	cargo clippy --all

tsc:
	tsc --project tsconfig.json

run:
	cargo run
