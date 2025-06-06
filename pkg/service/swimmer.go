package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid" // For generating unique IDs
	"github.com/l4rma/swim-api/pkg/models"
	"github.com/l4rma/swim-api/pkg/repository"
)

// SwimmerService defines the business logic for swimmers.
type SwimmerService interface {
	AddSwimmer(ctx context.Context, name string, age int) (*models.Swimmer, error)
	GetSwimmerById(ctx context.Context, swimmerID string) (*models.SwimmerSummary, error)
	UpdateSwimmer(ctx context.Context, id string, name string, age string) error
	DeleteSwimmer(ctx context.Context, id string) error
	ListSwimmers(ctx context.Context) ([]models.Swimmer, error)
}

var (
	r repository.SwimmerAndSessionRepository
)

type swimmerServiceImpl struct{}

func NewSwimmerService(repo repository.SwimmerAndSessionRepository) SwimmerService {
	r = repo
	return &swimmerServiceImpl{}
}

// AddSwimmer creates a new swimmer and adds it to the repository.
func (s *swimmerServiceImpl) AddSwimmer(ctx context.Context, name string, age int) (*models.Swimmer, error) {
	swimmer := models.Swimmer{
		ID:        uuid.NewString(),
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	if err := r.AddSwimmer(ctx, swimmer); err != nil {
		log.Printf("Failed to add swimmer: %v", err)
		return nil, fmt.Errorf("failed to add swimmer: %w", err)
	}
	return &swimmer, nil
}

func (s *swimmerServiceImpl) GetSwimmerById(ctx context.Context, swimmerID string) (*models.SwimmerSummary, error) {
	swimmer, err := r.GetSwimmerProfile(ctx, swimmerID)
	log.Printf("Service: Found swimmer: %+v", swimmer)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve swimmer profile: %w", err)
	}

	sessionSummary, err := r.SummarizeSwimmerSessions(ctx, swimmerID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve swimmer sessions: %w", err)
	}
	log.Printf("Service: Found session summary: %+v", sessionSummary)

	// Step 3: Combine the results into a SwimmerSummary
	return &models.SwimmerSummary{
		Swimmer:       *swimmer,
		TotalSessions: sessionSummary.TotalSessions,
		TotalDistance: sessionSummary.TotalDistance,
		TotalTime:     sessionSummary.TotalTime,
	}, nil
}

func (s *swimmerServiceImpl) UpdateSwimmer(ctx context.Context, id string, name string, age string) error {
	swimmer, err := r.GetSwimmerProfile(ctx, id)
	if err != nil {
		log.Printf("Failed to retrieve swimmer profile: %v", err)
		return fmt.Errorf("failed to retrieve swimmer profile: %w", err)
	}

	if name != "" {
		swimmer.Name = name
	}
	if age != "" {
		swimmer.Age, err = strconv.Atoi(age)
		if err != nil {
			log.Printf("Failed to parse age: %v", err)
			return fmt.Errorf("failed to parse age: %w", err)
		}
	}

	err = r.UpdateSwimmer(ctx, *swimmer)
	if err != nil {
		log.Printf("Failed to update swimmer: %v", err)
		return fmt.Errorf("failed to update swimmer: %w", err)
	}

	return nil
}

// DeleteSwimmer deactivates a swimmer by their ID.
func (s *swimmerServiceImpl) DeleteSwimmer(ctx context.Context, id string) error {
	swimmer, err := r.GetSwimmerProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("swimmer not found: %w", err)
	}
	swimmer.IsActive = false
	return r.UpdateSwimmer(ctx, *swimmer)
}

// ListSwimmers lists all swimmers.
func (s *swimmerServiceImpl) ListSwimmers(ctx context.Context) ([]models.Swimmer, error) {
	return r.ListSwimmers(ctx)
}
