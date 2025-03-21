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
	var request struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := json.Unmarshal([]byte(req.Body), &request); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to decode json"}, nil
	}
	swimmer, err := swimmerService.AddSwimmer(ctx, request.Name, request.Age)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to add swimmer"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: swimmer.ToString()}, nil
}

func main() {
	lambda.Start(handler)
}
