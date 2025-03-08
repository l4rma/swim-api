BINARY_NAME=bootstrap
APP=cmd/main.go

build:
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o ${BINARY_NAME} ${APP}

run: build
	@./${BINARY_NAME}

ddb:
	@docker-compose up -d
	@aws dynamodb create-table \
	    --table-name SwimmersAndSessions \
	    --attribute-definitions \
	        AttributeName=PK,AttributeType=S \
	        AttributeName=SK,AttributeType=S \
	    --key-schema \
	        AttributeName=PK,KeyType=HASH \
	        AttributeName=SK,KeyType=RANGE \
	    --provisioned-throughput \
	        ReadCapacityUnits=5,WriteCapacityUnits=5 \
	    --table-class STANDARD \
		--endpoint-url http://localhost:8000

sam: build ddb
	@cd infra && sam build --hook-name terraform && sam local start-api --docker-network dynamodb-local

clean:
	@go clean
	@rm bootstrap
	@rm infra/lambda_function_payload.zip

