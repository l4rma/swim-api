BINARY_NAME=bootstrap
APP=cmd/main.go
APIPATH=cmd/api/v2

build:
	@echo "Building create swimmer lambda"
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/create/${BINARY_NAME} ${APIPATH}/swimmers/create/main.go
	@echo "Building list swimmer lambda"
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/list/${BINARY_NAME} ${APIPATH}/swimmers/list/main.go
	@echo "Building find swimmer lambda"
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/find/${BINARY_NAME} ${APIPATH}/swimmers/find/main.go
	@echo "Building update swimmer lambda"
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/update/${BINARY_NAME} ${APIPATH}/swimmers/update/main.go
	@echo "Building delete swimmer lambda"
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/delete/${BINARY_NAME} ${APIPATH}/swimmers/delete/main.go
	@echo "Building create session lambda"
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bin/sessions/create/${BINARY_NAME} ${APIPATH}/sessions/create/main.go
	@echo "Build complete"

run: build
	@./${BINARY_NAME}

docker:
	@docker-compose up -d

table: docker
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

sam: build table
	@cd infra && sam build --hook-name terraform && sam local start-api --docker-network dynamodb-local

clean:
	go clean
	rm -rf bin/*
	rm -f infra/*.zip

