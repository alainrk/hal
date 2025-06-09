package node

import (
	"context"
	"fmt"
	"hal/pkg/graph"
	"hal/pkg/model"
	"hal/pkg/state"
)

// ModelNodeConfig contains configuration for a model node
type ModelNodeConfig struct {
	Model        model.Model
	PromptKey    string
	OutputKey    string
	SystemPrompt string
	Temperature  float32
}

// NewModelNode creates a new model node
func NewModelNode(name string, config *ModelNodeConfig) *graph.Node {
	return graph.NewNode(name, graph.NodeTypeModel, func(ctx context.Context, s *state.State) (*state.State, error) {
		// Get prompt from state
		promptValue, exists := s.Get(config.PromptKey)
		if !exists {
			return nil, fmt.Errorf("prompt key %s not found in state", config.PromptKey)
		}

		prompt, ok := promptValue.(string)
		if !ok {
			return nil, fmt.Errorf("prompt must be a string")
		}

		// Generate response
		response, err := config.Model.Generate(ctx, prompt, &model.GenerateOptions{
			SystemPrompt: config.SystemPrompt,
			Temperature:  config.Temperature,
		})
		if err != nil {
			return nil, fmt.Errorf("model generation failed: %w", err)
		}

		// Update state with response
		newState := s.Clone()
		newState.Set(config.OutputKey, response.Content)
		newState.Set(fmt.Sprintf("%s_metadata", config.OutputKey), response.Metadata)

		return newState, nil
	})
}
