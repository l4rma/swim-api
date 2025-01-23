package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/l4rma/swim-api/pkg/models"
	"github.com/l4rma/swim-api/pkg/utils"
)

type SwimmerAndSessionRepository interface {
	AddSwimmer(ctx context.Context, swimmer models.Swimmer) error
	GetSwimmerProfile(ctx context.Context, swimmerID string) (*models.Swimmer, error)
	SummarizeSwimmerSessions(ctx context.Context, swimmerID string) (*models.SessionSummary, error)
	UpdateSwimmer(ctx context.Context, swimmer models.Swimmer) error
	ListSwimmers(ctx context.Context) ([]models.Swimmer, error)
	AddSession(ctx context.Context, session models.Session) error
}

// DynamoDBRepository manages DynamoDB interactions.
type DynamoDBRepository struct {
	client    *dynamodb.Client
	tableName string
}

// NewDynamoDBRepository initializes a new DynamoDB repository.
func NewDynamoDBRepository(tableName string) (SwimmerAndSessionRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &DynamoDBRepository{
		client:    dynamodb.NewFromConfig(cfg),
		tableName: tableName,
	}, nil
}

func (repo *DynamoDBRepository) AddSwimmer(ctx context.Context, swimmer models.Swimmer) error {
	log.Printf("Repo: Adding swimmer: %+v", swimmer)
	item := map[string]types.AttributeValue{
		"PK":        &types.AttributeValueMemberS{Value: fmt.Sprintf("SWIMMER#%s", swimmer.ID)},
		"SK":        &types.AttributeValueMemberS{Value: "PROFILE"},
		"Name":      &types.AttributeValueMemberS{Value: swimmer.Name},
		"Age":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", swimmer.Age)},
		"CreatedAt": &types.AttributeValueMemberS{Value: swimmer.CreatedAt.Format("2006-01-02T15:04:05Z")}, // ISO 8601 format
		"IsActive":  &types.AttributeValueMemberBOOL{Value: swimmer.IsActive},
	}

	// Perform the PutItem operation
	_, err := repo.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(repo.tableName),
		Item:      item,
	})
	if err != nil {
		log.Printf("Failed to add swimmer to DynamoDB: %v", err)
		return fmt.Errorf("failed to add swimmer to DynamoDB: %w", err)
	}

	return nil
}

func (repo *DynamoDBRepository) GetSwimmerProfile(ctx context.Context, swimmerID string) (*models.Swimmer, error) {
	// Define the key to get the swimmer profile
	profileKey := map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("SWIMMER#%s", swimmerID)},
		"SK": &types.AttributeValueMemberS{Value: "PROFILE"},
	}

	// Perform the GetItem operation
	profileResp, err := repo.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key:       profileKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch swimmer profile: %w", err)
	}

	if profileResp.Item == nil {
		return nil, fmt.Errorf("swimmer not found")
	}

	// Parse the swimmer profile
	swimmer := &models.Swimmer{
		ID:        swimmerID,
		Name:      profileResp.Item["Name"].(*types.AttributeValueMemberS).Value,
		Age:       utils.ParseInt(profileResp.Item["Age"]),
		CreatedAt: utils.ParseTime(profileResp.Item["CreatedAt"]),
		IsActive:  profileResp.Item["IsActive"].(*types.AttributeValueMemberBOOL).Value,
	}

	return swimmer, nil
}

func (repo *DynamoDBRepository) SummarizeSwimmerSessions(ctx context.Context, swimmerID string) (*models.SessionSummary, error) {
	// Query all sessions for the swimmer
	sessionResp, err := repo.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(repo.tableName),
		KeyConditionExpression: aws.String("PK = :pk AND begins_with(SK, :prefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":     &types.AttributeValueMemberS{Value: fmt.Sprintf("SWIMMER#%s", swimmerID)},
			":prefix": &types.AttributeValueMemberS{Value: "SESSION#"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query swimmer sessions: %w", err)
	}

	// Initialize summary variables
	totalSessions := 0
	totalDistance := 0
	totalTime := 0

	// Aggregate session data
	for _, item := range sessionResp.Items {
		totalSessions++
		totalDistance += utils.ParseInt(item["Distance"])
		totalTime += utils.ParseInt(item["Duration"])
	}

	return &models.SessionSummary{
		TotalSessions: totalSessions,
		TotalDistance: totalDistance,
		TotalTime:     totalTime,
	}, nil
}

func (repo *DynamoDBRepository) UpdateSwimmer(ctx context.Context, swimmer models.Swimmer) error {
	// Define the key for the swimmer profile
	profileKey := map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("SWIMMER#%s", swimmer.ID)},
		"SK": &types.AttributeValueMemberS{Value: "PROFILE"},
	}

	// Prepare the update expression
	updateExpression := "SET Name = :name, Age = :age, IsActive = :is_active"
	expressionAttributeValues := map[string]types.AttributeValue{
		":name":      &types.AttributeValueMemberS{Value: swimmer.Name},
		":age":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", swimmer.Age)},
		":is_active": &types.AttributeValueMemberBOOL{Value: swimmer.IsActive},
	}

	// Execute the UpdateItem operation
	_, err := repo.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(repo.tableName),
		Key:                       profileKey,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: expressionAttributeValues,
	})
	if err != nil {
		return fmt.Errorf("failed to update swimmer: %w", err)
	}

	return nil
}

func (repo *DynamoDBRepository) ListSwimmers(ctx context.Context) ([]models.Swimmer, error) {
	// Perform a Scan operation to fetch all swimmer profiles
	resp, err := repo.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(repo.tableName),
		FilterExpression: aws.String("begins_with(PK, :prefix) AND SK = :profile"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":prefix":  &types.AttributeValueMemberS{Value: "SWIMMER#"},
			":profile": &types.AttributeValueMemberS{Value: "PROFILE"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch swimmers: %w", err)
	}

	// Parse the results into a slice of Swimmer objects
	swimmers := []models.Swimmer{}
	for _, item := range resp.Items {
		swimmer := models.Swimmer{
			ID:        utils.ParseString(item["PK"]),
			Name:      utils.ParseString(item["Name"]),
			Age:       utils.ParseInt(item["Age"]),
			CreatedAt: utils.ParseTime(item["CreatedAt"]),
			IsActive:  utils.ParseBool(item["IsActive"]),
		}
		// Extract the ID from the PK (remove "SWIMMER#" prefix)
		swimmer.ID = swimmer.ID[len("SWIMMER#"):]
		swimmers = append(swimmers, swimmer)
	}

	return swimmers, nil
}

func (repo *DynamoDBRepository) AddSession(ctx context.Context, session models.Session) error {
	log.Printf("Adding session: %+v", session)
	_, err := repo.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(repo.tableName),
		Item: map[string]types.AttributeValue{
			"PK":        &types.AttributeValueMemberS{Value: fmt.Sprintf("SWIMMER#%s", session.SwimmerID)},
			"SK":        &types.AttributeValueMemberS{Value: fmt.Sprintf("SESSION#%s", session.ID)},
			"Date":      &types.AttributeValueMemberS{Value: session.Date.Format("2006-01-02")},
			"Distance":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", session.Distance)},
			"Duration":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", int(session.Duration.Minutes()))},
			"Intensity": &types.AttributeValueMemberS{Value: session.Intensity},
			"Style":     &types.AttributeValueMemberS{Value: session.Style},
			"Notes":     &types.AttributeValueMemberS{Value: session.Notes},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to add session: %w", err)
	}

	return nil
}
