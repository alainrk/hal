package graph

import (
	"context"
	"hal/pkg/state"

	"github.com/google/uuid"
)

// NodeFunc is the function signature for node execution
type NodeFunc func(ctx context.Context, state *state.State) (*state.State, error)

// Node represents a node in the graph
type Node struct {
	ID          string
	Name        string
	Type        NodeType
	Execute     NodeFunc
	Config      map[string]any
	RetryPolicy *RetryPolicy
}

// NodeType represents the type of node
type NodeType string

const (
	NodeTypeModel    NodeType = "model"
	NodeTypeFunction NodeType = "function"
	NodeTypeRouter   NodeType = "router"
	NodeTypeStart    NodeType = "start"
	NodeTypeEnd      NodeType = "end"
)

// RetryPolicy defines retry behavior for a node
type RetryPolicy struct {
	MaxAttempts int
	BackoffMs   int
}

// NewNode creates a new node
func NewNode(name string, nodeType NodeType, fn NodeFunc) *Node {
	return &Node{
		ID:      uuid.New().String(),
		Name:    name,
		Type:    nodeType,
		Execute: fn,
		Config:  make(map[string]any),
	}
}
