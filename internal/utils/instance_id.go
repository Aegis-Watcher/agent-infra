package utils

import (
	"log"

	"github.com/google/uuid"
)

// GenerateInstanceID generates a unique instance ID for the agent metrics.
// It uses the google/uuid library to generate a random UUID.
func GenerateInstanceID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Failed to generate instance ID: %v", err)
	}
	return id.String()
}
