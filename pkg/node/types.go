package node

import (
	"hal/pkg/graph"
)

// NodeBuilder provides a fluent interface for building nodes
type NodeBuilder struct {
	node *graph.Node
}

// NewBuilder creates a new node builder
func NewBuilder(name string) *NodeBuilder {
	return &NodeBuilder{
		node: &graph.Node{
			Name:   name,
			Config: make(map[string]any),
		},
	}
}

// WithType sets the node type
func (b *NodeBuilder) WithType(nodeType graph.NodeType) *NodeBuilder {
	b.node.Type = nodeType
	return b
}

// WithRetry sets the retry policy
func (b *NodeBuilder) WithRetry(maxAttempts int, backoffMs int) *NodeBuilder {
	b.node.RetryPolicy = &graph.RetryPolicy{
		MaxAttempts: maxAttempts,
		BackoffMs:   backoffMs,
	}
	return b
}

// WithConfig adds configuration
func (b *NodeBuilder) WithConfig(key string, value any) *NodeBuilder {
	b.node.Config[key] = value
	return b
}

// Build returns the constructed node
func (b *NodeBuilder) Build() *graph.Node {
	return b.node
}
