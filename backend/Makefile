.PHONY: codegen
codegen:
	rm -rf ./codegen/src/main/java/ee/fakeplastictrees/blog/codegen/*
	docker run --rm \
			-v ./spec:/workdir/spec:ro \
			-v ./codegen/src/main/java/eu/raweceek/codegen/models:/workdir/out/src/main/java/eu/raweceek/codegen/models \
			-v ./codegen/src/main/java/eu/raweceek/codegen/api:/workdir/out/src/main/java/eu/raweceek/codegen/api \
			openapitools/openapi-generator-cli generate \
			-c /workdir/spec/swagger-config.yaml \
			-i /workdir/spec/swagger.yaml \
			-g spring \
			-o /workdir/out

.PHONY: build
build:
	docker build -t rawe-ceek-backend:latest .
	docker save rawe-ceek-backend:latest > rawe-ceek-backend.tar

.PHONY: push
push:
	rsync rawe-ceek-backend.tar $(RAWE_CEEK_BACKEND_PWD)
