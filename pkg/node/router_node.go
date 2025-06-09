package node

import (
	"context"
	"hal/pkg/graph"
	"hal/pkg/state"
)

// RouterNodeConfig contains configuration for a router node
type RouterNodeConfig struct {
	DecisionKey string
	Routes      map[string]string // value -> nodeID mapping
}

// NewRouterNode creates a new router node
func NewRouterNode(name string, config *RouterNodeConfig) *graph.Node {
	return graph.NewNode(name, graph.NodeTypeRouter, func(ctx context.Context, s *state.State) (*state.State, error) {
		// Router nodes don't modify state, they just help with routing
		// The actual routing logic is handled by conditional edges
		return s, nil
	})
}
