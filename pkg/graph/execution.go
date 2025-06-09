package graph

import (
	"context"
	"fmt"
	"hal/pkg/state"
	"time"
)

// ExecutionResult represents the result of graph execution
type ExecutionResult struct {
	State     *state.State
	StartTime time.Time
	EndTime   time.Time
	NodesRun  []string
	Error     error
}

// Execute runs the graph with the given initial state
func (g *Graph) Execute(ctx context.Context, initialState *state.State) (*ExecutionResult, error) {
	result := &ExecutionResult{
		StartTime: time.Now(),
		State:     initialState,
		NodesRun:  make([]string, 0),
	}

	currentNodeID := g.entryPoint

	for currentNodeID != "" {
		select {
		case <-ctx.Done():
			result.Error = ctx.Err()
			result.EndTime = time.Now()
			return result, ctx.Err()
		default:
		}

		node, exists := g.nodes[currentNodeID]
		if !exists {
			result.Error = fmt.Errorf("node %s not found", currentNodeID)
			result.EndTime = time.Now()
			return result, result.Error
		}

		// Execute node
		newState, err := g.executeNode(ctx, node, result.State)
		if err != nil {
			result.Error = err
			result.EndTime = time.Now()
			return result, err
		}

		result.State = newState
		result.NodesRun = append(result.NodesRun, currentNodeID)

		// Find next node
		nextNodeID := g.findNextNode(currentNodeID, result.State)
		currentNodeID = nextNodeID
	}

	result.EndTime = time.Now()
	return result, nil
}

// executeNode executes a single node with retry logic
func (g *Graph) executeNode(ctx context.Context, node *Node, state *state.State) (*state.State, error) {
	retryPolicy := node.RetryPolicy
	if retryPolicy == nil {
		retryPolicy = &RetryPolicy{MaxAttempts: 1}
	}

	var lastErr error
	for attempt := 0; attempt < retryPolicy.MaxAttempts; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(retryPolicy.BackoffMs) * time.Millisecond)
		}

		newState, err := node.Execute(ctx, state)
		if err == nil {
			return newState, nil
		}

		lastErr = err
	}

	return nil, fmt.Errorf("node %s failed after %d attempts: %w",
		node.Name, retryPolicy.MaxAttempts, lastErr)
}

// findNextNode determines the next node based on edges and conditions
func (g *Graph) findNextNode(currentNodeID string, state *state.State) string {
	edges := g.edges[currentNodeID]

	for _, edge := range edges {
		if edge.Condition(state) {
			return edge.To
		}
	}

	return "" // No matching edge found
}
