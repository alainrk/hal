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

// Generate implements Model.Generate
func (m *AnthropicModel) Generate(ctx context.Context, prompt string, options *GenerateOptions) (*Response, error) {
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
