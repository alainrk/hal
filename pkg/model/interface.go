package model

import (
	"context"
)

// Model defines the interface for language models
type Model interface {
	// Generate generates a response based on the prompt
	Generate(ctx context.Context, prompt string, options *GenerateOptions) (*Response, error)

	// GetName returns the model name
	GetName() string
}

// GenerateOptions contains options for generation
type GenerateOptions struct {
	Temperature   float32
	MaxTokens     int
	TopP          float32
	StopSequences []string
	SystemPrompt  string
}

// Response represents a model response
type Response struct {
	Content      string
	TokensUsed   int
	FinishReason string
	Metadata     map[string]interface{}
}

// Message represents a chat message
type Message struct {
	Role    string
	Content string
}
