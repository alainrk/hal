package node

import (
	"context"
	"hal/pkg/graph"
	"hal/pkg/state"
)

// FunctionNodeConfig contains configuration for a function node
type FunctionNodeConfig struct {
	Function  func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error)
	InputKeys []string
}

// NewFunctionNode creates a new function node
func NewFunctionNode(name string, config *FunctionNodeConfig) *graph.Node {
	return graph.NewNode(name, graph.NodeTypeFunction, func(ctx context.Context, s *state.State) (*state.State, error) {
		// Gather inputs
		inputs := make(map[string]interface{})
		for _, key := range config.InputKeys {
			if val, exists := s.Get(key); exists {
				inputs[key] = val
			}
		}

		// Execute function
		outputs, err := config.Function(ctx, inputs)
		if err != nil {
			return nil, err
		}

		// Update state
		newState := s.Clone()
		newState.Update(outputs)

		return newState, nil
	})
}
