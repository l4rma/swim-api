BINARY_NAME=bootstrap
APP=cmd/main.go

build:
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o ${BINARY_NAME} ${APP}

run: build
	@./${BINARY_NAME}

sam: build
	@cd infra && sam build --hook-name terraform && sam local start-api --docker-network dynamodb-local

clean:
	@go clean
	@rm bootstrap
	@rm infra/lambda_function_payload.zip

