codegen_run:
	rm -rf ./codegen/src/main/java/eu/raweceek/codegen/*
	docker run --name raweceek-backend-codegen \
    		-v ./spec:/workdir/spec:ro \
            openapitools/openapi-generator-cli generate \
    		-c /workdir/spec/swagger-config.yaml \
            -i /workdir/spec/swagger.yaml \
            -g spring \
            -o /workdir/out
	docker cp raweceek-backend-codegen:/workdir/out/src/main/java/eu/raweceek/codegen/models/ ./codegen/src/main/java/eu/raweceek/codegen/models/
	docker cp raweceek-backend-codegen:/workdir/out/src/main/java/eu/raweceek/codegen/api/ ./codegen/src/main/java/eu/raweceek/codegen/api/

codegen_cleanup:
	docker rm raweceek-backend-codegen

.PHONY: codegen
codegen:
	$(MAKE) codegen_run || exit 0
	$(MAKE) codegen_cleanup
