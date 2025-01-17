package service

import (
	"fmt"
	"time"

	"github.com/google/uuid" // For generating unique IDs
	"github.com/l4rma/swim-api/pkg/models"
	"github.com/l4rma/swim-api/pkg/repository/inmemory"
)

// SwimmerService defines the business logic for swimmers.
type SwimmerService interface {
	AddSwimmer(name string, age int) (*models.Swimmer, error)
	GetSwimmerByID(id string) (*models.Swimmer, error)
	UpdateSwimmer(id string, name string, age int) error
	DeleteSwimmer(id string) error
	ListSwimmers() ([]models.Swimmer, error)
}

var (
	repository inmemory.SwimmerRepository
)

// swimmerServiceImpl is the implementation of SwimmerService.
type swimmerServiceImpl struct{}

// NewSwimmerService creates a new SwimmerService.
func NewSwimmerService(repo inmemory.SwimmerRepository) SwimmerService {
	repository = repo
	return &swimmerServiceImpl{}
}

// AddSwimmer creates a new swimmer and adds it to the repository.
func (s *swimmerServiceImpl) AddSwimmer(name string, age int) (*models.Swimmer, error) {
	swimmer := models.Swimmer{
		ID:        uuid.NewString(),
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
	if err := repository.AddSwimmer(swimmer); err != nil {
		return nil, fmt.Errorf("failed to add swimmer: %w", err)
	}
	return &swimmer, nil
}

// GetSwimmerByID retrieves a swimmer by their ID.
func (s *swimmerServiceImpl) GetSwimmerByID(id string) (*models.Swimmer, error) {
	return repository.GetSwimmerByID(id)
}

// UpdateSwimmer updates an existing swimmer's details.
func (s *swimmerServiceImpl) UpdateSwimmer(id string, name string, age int) error {
	swimmer, err := repository.GetSwimmerByID(id)
	if err != nil {
		return fmt.Errorf("swimmer not found: %w", err)
	}
	swimmer.Name = name
	swimmer.Age = age
	return repository.UpdateSwimmer(*swimmer)
}

// DeleteSwimmer deactivates a swimmer by their ID.
func (s *swimmerServiceImpl) DeleteSwimmer(id string) error {
	swimmer, err := repository.GetSwimmerByID(id)
	if err != nil {
		return fmt.Errorf("swimmer not found: %w", err)
	}
	swimmer.IsActive = false
	return repository.UpdateSwimmer(*swimmer)
}

// ListSwimmers lists all swimmers.
func (s *swimmerServiceImpl) ListSwimmers() ([]models.Swimmer, error) {
	return repository.ListSwimmers()
}
