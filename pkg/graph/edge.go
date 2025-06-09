package graph

import (
	"hal/pkg/state"

	"github.com/google/uuid"
)

// Edge represents a connection between nodes
type Edge struct {
	ID        string
	From      string
	To        string
	Condition EdgeCondition
	Weight    float64
}

// EdgeCondition is a function that determines if an edge should be traversed
type EdgeCondition func(state *state.State) bool

// NewEdge creates a new edge
func NewEdge(from, to string) *Edge {
	return &Edge{
		ID:     uuid.New().String(),
		From:   from,
		To:     to,
		Weight: 1.0,
		Condition: func(s *state.State) bool {
			return true // Default: always traverse
		},
	}
}

// NewConditionalEdge creates an edge with a condition
func NewConditionalEdge(from, to string, condition EdgeCondition) *Edge {
	edge := NewEdge(from, to)
	edge.Condition = condition
	return edge
}
