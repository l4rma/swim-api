BINARY_NAME=bootstrap
APP=cmd/main.go

build:
	@GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o ${BINARY_NAME} ${APP}

run: build
	@./${BINARY_NAME}

clean:
	@go clean
	@rm bootstrap

