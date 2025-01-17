package service

import (
	"fmt"
	"time"

	"github.com/google/uuid" // For generating unique IDs
	"github.com/l4rma/swim-api/pkg/models"
	"github.com/l4rma/swim-api/pkg/repository/inmemory"
)

// SessionService defines the business logic for sessions.
type SessionService interface {
	AddSession(swimmerID string, date time.Time, distance int, minutes int, intensity, style, notes string) (*models.Session, error)
	GetSessionByID(id string) (*models.Session, error)
	GetSessionsBySwimmerID(swimmerID string) ([]models.Session, error)
	UpdateSession(id string, date time.Time, distance int, duration time.Duration, intensity, style, notes string) error
	DeleteSession(id string) error
	ListSessions() ([]models.Session, error)
}

var (
	sessionRepository inmemory.SessionRepository
)

// sessionServiceImpl is the implementation of SessionService.
type sessionServiceImpl struct{}

// NewSessionService creates a new SessionService.
func NewSessionService(repo inmemory.SessionRepository) SessionService {
	sessionRepository = repo
	return &sessionServiceImpl{}
}

// AddSession creates a new session and adds it to the repository.
func (s *sessionServiceImpl) AddSession(swimmerID string, date time.Time, distance int, minutes int, intensity, style, notes string) (*models.Session, error) {
	duration := time.Duration(minutes) * time.Minute
	session := models.Session{
		ID:        uuid.NewString(),
		SwimmerID: swimmerID,
		Date:      date,
		Distance:  distance,
		Duration:  duration,
		Intensity: intensity,
		Style:     style,
		Notes:     notes,
	}
	if err := sessionRepository.AddSession(session); err != nil {
		return nil, fmt.Errorf("failed to add session: %w", err)
	}
	return &session, nil
}

// GetSessionByID retrieves a session by its ID.
func (s *sessionServiceImpl) GetSessionByID(id string) (*models.Session, error) {
	return sessionRepository.GetSessionByID(id)
}

// GetSessionsBySwimmerID retrieves all sessions for a specific swimmer.
func (s *sessionServiceImpl) GetSessionsBySwimmerID(swimmerID string) ([]models.Session, error) {
	return sessionRepository.GetSessionsBySwimmerID(swimmerID)
}

// UpdateSession updates an existing session's details.
func (s *sessionServiceImpl) UpdateSession(id string, date time.Time, distance int, duration time.Duration, intensity, style, notes string) error {
	session, err := sessionRepository.GetSessionByID(id)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}
	session.Date = date
	session.Distance = distance
	session.Duration = duration
	session.Intensity = intensity
	session.Style = style
	session.Notes = notes
	return sessionRepository.UpdateSession(*session)
}

// DeleteSession removes a session by its ID.
func (s *sessionServiceImpl) DeleteSession(id string) error {
	return sessionRepository.DeleteSession(id)
}

// ListSessions lists all sessions.
func (s *sessionServiceImpl) ListSessions() ([]models.Session, error) {
	return sessionRepository.ListSessions()
}
