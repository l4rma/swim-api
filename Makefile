BINARY_NAME=bootstrap
APP=cmd/main.go
APIPATH=cmd/api/v2

build:
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o ${BINARY_NAME} ${APP}
	@GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -o bin/create/bootstrap ${APIPATH}/swimmers/create/main.go
	@GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -o bin/list/bootstrap ${APIPATH}/swimmers/list/main.go
	@GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -o bin/update/bootstrap ${APIPATH}/swimmers/update/main.go

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
	@rm -r bin/*
	@rm infra/*.zip
	@rm bootstrap

