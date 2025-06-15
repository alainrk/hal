package graph

import (
	"context"
)

// Runnable is the fundamental unit of work in the graph.
// It is a generic interface that operates on any user-defined state type 'S'.
type Runnable[S any] interface {
	// Run executes the node's logic. It receives the current state
	// and returns the potentially modified state.
	Run(ctx context.Context, state S) (S, error)
}

// SimpleNode is the simplest node we give to the user.
// It just implements the Runnable interface.
type SimpleNode[S any] func(ctx context.Context, state S) (S, error)

// Run simply calls the underlying function, without doing anything else, satisfying the Runnable interface.
//
// Example:
//
//	type AgentState struct {
//		Messages []hal.ChatMessage
//	}
//
//	func callModelNode(ctx context.Context, state AgentState) (AgentState, error) {
//		// ... implementation
//		return state, nil
//	}
//
//	modelNode := hal.SimpleNode[AgentState](callModelNode)
//
//	initialState := AgentState{
//		Messages: []hal.ChatMessage{Role: "user", Content: "Hello, world!"},
//	}
//
//	finalState, err := modelNode.Run(context.Background(), initialState)
func (f SimpleNode[S]) Run(ctx context.Context, state S) (S, error) {
	return f(ctx, state)
}
