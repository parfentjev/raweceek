-include .env
export

.PHONY: fmt lint test tsc run

fmt:
	cargo fmt

lint:
	cargo clippy --all

test:
	cargo test

tsc:
	tsc --project tsconfig.json

run:
	cargo run
