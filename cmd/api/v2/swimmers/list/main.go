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
	dbUrl := ""
	repo, err := repository.NewDynamoDBRepository(dbUrl, "SwimmersAndSessions")
	if err != nil {
		log.Fatalf("failed to create DynamoDB repository: %v", err)
	}
	swimmerService = service.NewSwimmerService(repo)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	swimmers, err := swimmerService.ListSwimmers(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to list swimmers"}, nil
	}

	marshalledSwimmers, err := json.Marshal(swimmers)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(marshalledSwimmers)}, nil
}

func main() {
	lambda.Start(handler)
}
