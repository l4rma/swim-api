package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/l4rma/swim-api/pkg/repository"
	"github.com/l4rma/swim-api/pkg/service"
)

var (
	swimmerService service.SwimmerService
	sessionService service.SessionService
)

func init() {
	dbUrl := "http://dynamodb:8000" // Local DynamoDB
	//dbUrl := "" // AWS DynamoDB

	repo, err := repository.NewDynamoDBRepository(dbUrl, "SwimmersAndSessions")
	if err != nil {
		log.Fatalf("failed to create DynamoDB repository: %v", err)
	}
	swimmerService = service.NewSwimmerService(repo)
	sessionService = service.NewSessionService(repo)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var request struct {
		SwimmerID string `json:"swimmer_id"`
		Date      string `json:"date"`
		Distance  int    `json:"distance"`
		Duration  int    `json:"duration"`
		Intensity string `json:"intensity"`
		Style     string `json:"style"`
		Notes     string `json:"notes"`
	}

	if err := json.Unmarshal([]byte(req.Body), &request); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to decode json"}, nil
	}
	parsedDate, err := time.Parse("2006.01.02", request.Date)
	if err != nil {
		log.Printf("Failed to parse date: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Invalid date format, expected YYYY.MM.DD"}, nil
	}

	session, err := sessionService.AddSession(ctx, request.SwimmerID, parsedDate, request.Distance, request.Duration, request.Intensity, request.Style, request.Notes)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to add swimmer"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: session.ToString()}, nil
}

func main() {
	lambda.Start(handler)
}
