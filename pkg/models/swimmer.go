package models

import (
	"encoding/json"
	"log"
	"time"
)

type Swimmer struct {
	ID        string    `json:"id"`         // Unique identifier for the swimmer
	Name      string    `json:"name"`       // Swimmer's full name
	Age       int       `json:"age"`        // Swimmer's age
	CreatedAt time.Time `json:"created_at"` // Timestamp when the swimmer was added
	IsActive  bool      `json:"is_active"`  // Indicates whether the swimmer is active
}

func (s Swimmer) ToString() string {
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Error marshalling swimmer to JSON %v", err)
	}
	return string(jsonData)
}
