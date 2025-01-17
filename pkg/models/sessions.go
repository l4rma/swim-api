package models

import "time"

type Session struct {
	ID        string        `json:"id"`         // Unique identifier for the session
	SwimmerID string        `json:"swimmer_id"` // Foreign key linking to Swimmer.ID
	Date      time.Time     `json:"date"`       // Date and time of the session
	Distance  int           `json:"distance"`   // Total distance swam in meters
	Duration  time.Duration `json:"duration"`   // Total duration of the session
	Intensity string        `json:"intensity"`  // Intensity level (e.g., "low", "moderate", "high")
	Style     string        `json:"style"`      // Swimming style (e.g., "freestyle", "butterfly", "mixed")
	Notes     string        `json:"notes"`      // Additional notes about the session
}
