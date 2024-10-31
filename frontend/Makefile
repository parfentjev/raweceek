.PHONY: start
start:
	BROWSER=none npm start

.PHONY: codegen
codegen:
	rm -rf ./src/api/codegen
	docker run --rm \
		-v ./../backend/spec:/workdir/spec:ro \
		-v ./src/api/codegen:/workdir/out \
		openapitools/openapi-generator-cli generate \
		-i /workdir/spec/swagger.yaml \
		-g typescript-fetch \
		-o /workdir/out \
		--additional-properties hooks=true

.PHONY: build
build:
	NODE_ENV="production" REACT_APP_SERVICE_URL="https://raweceek.eu/api" npm run build

.PHONY: push
push:
	rsync -r build/* $(RAWE_CEEK_FRONTEND_PWD)
