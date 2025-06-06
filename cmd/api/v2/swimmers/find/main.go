package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/l4rma/swim-api/pkg/repository"
	"github.com/l4rma/swim-api/pkg/service"
)

var (
	swimmerService service.SwimmerService
)

func init() {
	dbUrl := "http://dynamodb:8000" // Local DynamoDB
	//dbUrl := "" // AWS DynamoDB

	repo, err := repository.NewDynamoDBRepository(dbUrl, "SwimmersAndSessions")
	if err != nil {
		log.Fatalf("failed to create DynamoDB repository: %v", err)
	}
	swimmerService = service.NewSwimmerService(repo)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	if id == "" {
		log.Println("Error: Missing swimmer ID")
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to list swimmers"}, nil
	}
	swimmer, err := swimmerService.GetSwimmerById(ctx, id)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to find swimmer"}, nil
	}

	marshalledSwimmer, err := json.Marshal(swimmer)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(marshalledSwimmer)}, nil
}

func main() {
	lambda.Start(handler)
}
