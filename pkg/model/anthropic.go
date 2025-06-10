package model

import (
	"context"
	// Anthropic client implementation would go here
)

// AnthropicModel implements the Model interface for Anthropic Claude
type AnthropicModel struct {
	apiKey    string
	modelName string
}

// NewAnthropicModel creates a new Anthropic model
func NewAnthropicModel(apiKey string, modelName string) *AnthropicModel {
	return &AnthropicModel{
		apiKey:    apiKey,
		modelName: modelName,
	}
}

// Invoke implements Model.Generate
func (m *AnthropicModel) Invoke(ctx context.Context, prompt string, options *InvokeOptions) (*Response, error) {
	// TODO: Implement Anthropic API call
	return &Response{
		Content: "Anthropic implementation pending",
		Metadata: map[string]any{
			"model": m.modelName,
		},
	}, nil
}

// GetName returns the model name
func (m *AnthropicModel) GetName() string {
	return m.modelName
}
