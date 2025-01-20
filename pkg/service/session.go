package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid" // For generating unique IDs
	"github.com/l4rma/swim-api/pkg/models"
	"github.com/l4rma/swim-api/pkg/repository"
)

// SessionService defines the business logic for sessions.
type SessionService interface {
	AddSession(ctx context.Context, swimmerID string, date time.Time, distance int, minutes int, intensity, style, notes string) (*models.Session, error)
}

// sessionServiceImpl is the implementation of SessionService.
type sessionServiceImpl struct{}

// NewSessionService creates a new SessionService.
func NewSessionService(repo repository.SwimmerAndSessionRepository) SessionService {
	r = repo
	return &sessionServiceImpl{}
}

// AddSession creates a new session and adds it to the repository.
func (s *sessionServiceImpl) AddSession(ctx context.Context, swimmerID string, date time.Time, distance int, minutes int, intensity, style, notes string) (*models.Session, error) {
	log.Printf("Adding session for swimmer %s", swimmerID)
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
	if err := r.AddSession(ctx, session); err != nil {
		return nil, fmt.Errorf("failed to add session: %w", err)
	}
	return &session, nil
}
