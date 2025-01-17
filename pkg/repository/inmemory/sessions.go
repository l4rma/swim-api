package inmemory

import (
	"errors"
	"sync"

	"github.com/l4rma/swim-api/pkg/models"
)

// SessionRepository defines CRUD operations for sessions.
type SessionRepository interface {
	AddSession(session models.Session) error
	GetSessionByID(id string) (*models.Session, error)
	GetSessionsBySwimmerID(swimmerID string) ([]models.Session, error)
	UpdateSession(session models.Session) error
	DeleteSession(id string) error
	ListSessions() ([]models.Session, error)
}

// InMemorySessionRepository is an in-memory implementation of SessionRepository.
type InMemorySessionRepository struct {
	mu       sync.RWMutex
	sessions []models.Session
}

// NewInMemorySessionRepository creates a new in-memory session repository.
func NewInMemorySessionRepository() *InMemorySessionRepository {
	return &InMemorySessionRepository{
		sessions: make([]models.Session, 0),
	}
}

// AddSession adds a session to the repository.
func (r *InMemorySessionRepository) AddSession(session models.Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sessions = append(r.sessions, session)
	return nil
}

// GetSessionByID retrieves a session by its ID.
func (r *InMemorySessionRepository) GetSessionByID(id string) (*models.Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, session := range r.sessions {
		if session.ID == id {
			return &session, nil
		}
	}
	return nil, errors.New("session not found")
}

// GetSessionsBySwimmerID retrieves all sessions for a specific swimmer.
func (r *InMemorySessionRepository) GetSessionsBySwimmerID(swimmerID string) ([]models.Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []models.Session
	for _, session := range r.sessions {
		if session.SwimmerID == swimmerID {
			result = append(result, session)
		}
	}
	return result, nil
}

// UpdateSession updates an existing session's details.
func (r *InMemorySessionRepository) UpdateSession(session models.Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, s := range r.sessions {
		if s.ID == session.ID {
			r.sessions[i] = session
			return nil
		}
	}
	return errors.New("session not found")
}

// DeleteSession removes a session by its ID.
func (r *InMemorySessionRepository) DeleteSession(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, session := range r.sessions {
		if session.ID == id {
			r.sessions = append(r.sessions[:i], r.sessions[i+1:]...)
			return nil
		}
	}
	return errors.New("session not found")
}

// ListSessions returns all sessions in the repository.
func (r *InMemorySessionRepository) ListSessions() ([]models.Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.sessions, nil
}
