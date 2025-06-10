package utils

import (
	"github.com/google/uuid"
)

// GenerateID generates a unique ID
func GenerateID() string {
	return uuid.New().String()
}

// GenerateShortID generates a shorter unique ID
func GenerateShortID() string {
	id := uuid.New()
	return id.String()[:8]
}
