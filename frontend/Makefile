.PHONY: start
start:
	BROWSER=none npm start

codegen_run:
	rm -rf ./src/api/codegen
	docker run --name raweceek-frontend-codegen \
    		-v ./../backend/spec:/workdir/spec:ro \
            openapitools/openapi-generator-cli generate \
    		-c /workdir/spec/swagger-config.yaml \
            -i /workdir/spec/swagger.yaml \
            -g typescript-fetch \
            -o /workdir/codegen \
			--additional-properties hooks=true
	docker cp raweceek-frontend-codegen:/workdir/codegen/ ./src/api/

codegen_cleanup:
	docker rm raweceek-frontend-codegen

.PHONY: codegen
codegen:
	$(MAKE) codegen_run || exit 0
	$(MAKE) codegen_cleanup

.PHONY: build
build:
	NODE_ENV="production" REACT_APP_SERVICE_URL="https://raweceek.eu/api" npm run build

.PHONY: push
push:
	rsync -r build/* $(RAWE_CEEK_FRONTEND_PWD)
