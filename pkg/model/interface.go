package model

import (
	"context"
)

// Model defines the interface for language models
type Model interface {
	// Invoke generates a response based on the prompt
	Invoke(ctx context.Context, prompt string, options *InvokeOptions) (*Response, error)

	// GetName returns the model name
	GetName() string
}

// InvokeOptions contains options for generation
type InvokeOptions struct {
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
	Metadata     map[string]any
}

// Message represents a chat message
type Message struct {
	Role    string
	Content string
}
