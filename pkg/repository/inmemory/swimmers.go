package inmemory

import (
	"context"
	"errors"
	"sync"

	"github.com/l4rma/swim-api/pkg/models"
)

var (
	swimmer models.Swimmer
)

// SwimmerRepository defines CRUD operations for swimmers.
type SwimmerRepository interface {
	AddSwimmer(ctx context.Context, swimmer models.Swimmer) error
	GetSwimmerByID(id string) (*models.Swimmer, error)
	UpdateSwimmer(swimmer models.Swimmer) error
	DeleteSwimmer(id string) error
	ListSwimmers() ([]models.Swimmer, error)
}

// InMemorySwimmerRepository is an in-memory implementation of SwimmerRepository.
type InMemorySwimmerRepository struct {
	mu       sync.RWMutex
	swimmers []models.Swimmer
}

// NewInMemorySwimmerRepository creates a new in-memory swimmer repository.
func NewInMemorySwimmerRepository() *InMemorySwimmerRepository {
	return &InMemorySwimmerRepository{
		swimmers: make([]models.Swimmer, 0),
	}
}

// AddSwimmer adds a swimmer to the repository.
func (r *InMemorySwimmerRepository) AddSwimmer(ctx context.Context, swimmer models.Swimmer) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.swimmers = append(r.swimmers, swimmer)
	return nil
}

// GetSwimmerByID retrieves a swimmer by their ID.
func (r *InMemorySwimmerRepository) GetSwimmerByID(id string) (*models.Swimmer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, swimmer := range r.swimmers {
		if swimmer.ID == id {
			return &swimmer, nil
		}
	}
	return nil, errors.New("swimmer not found")
}

// UpdateSwimmer updates an existing swimmer's details.
func (r *InMemorySwimmerRepository) UpdateSwimmer(swimmer models.Swimmer) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, s := range r.swimmers {
		if s.ID == swimmer.ID {
			r.swimmers[i] = swimmer
			return nil
		}
	}
	return errors.New("swimmer not found")
}

// DeleteSwimmer removes a swimmer by their ID.
func (r *InMemorySwimmerRepository) DeleteSwimmer(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, swimmer := range r.swimmers {
		if swimmer.ID == id {
			r.swimmers = append(r.swimmers[:i], r.swimmers[i+1:]...)
			return nil
		}
	}
	return errors.New("swimmer not found")
}

// ListSwimmers returns all swimmers in the repository.
func (r *InMemorySwimmerRepository) ListSwimmers() ([]models.Swimmer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.swimmers, nil
}
