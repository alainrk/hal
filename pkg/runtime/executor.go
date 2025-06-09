package runtime

import (
	"context"
	"fmt"
	"hal/pkg/graph"
	"hal/pkg/state"
	"sync"
)

// Executor manages graph execution
type Executor struct {
	mu         sync.Mutex
	graphs     map[string]*graph.Graph
	executions map[string]*ExecutionContext
}

// ExecutionContext tracks an ongoing execution
type ExecutionContext struct {
	ID     string
	Graph  *graph.Graph
	State  *state.State
	Cancel context.CancelFunc
	Status ExecutionStatus
}

// ExecutionStatus represents the status of an execution
type ExecutionStatus string

const (
	StatusRunning   ExecutionStatus = "running"
	StatusCompleted ExecutionStatus = "completed"
	StatusFailed    ExecutionStatus = "failed"
	StatusCancelled ExecutionStatus = "cancelled"
)

// NewExecutor creates a new executor
func NewExecutor() *Executor {
	return &Executor{
		graphs:     make(map[string]*graph.Graph),
		executions: make(map[string]*ExecutionContext),
	}
}

// RegisterGraph registers a graph with the executor
func (e *Executor) RegisterGraph(name string, g *graph.Graph) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.graphs[name] = g
}

// Execute runs a graph by name
func (e *Executor) Execute(ctx context.Context, graphName string, initialState *state.State) (*graph.ExecutionResult, error) {
	e.mu.Lock()
	g, exists := e.graphs[graphName]
	e.mu.Unlock()

	if !exists {
		return nil, fmt.Errorf("graph %s not found", graphName)
	}

	return g.Execute(ctx, initialState)
}
