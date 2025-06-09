# HAL - A Graph-Based AI Framework for Go

HAL is a lightweight, efficient graph-based framework for building AI applications in Go.

## Features

- Graph-based execution model
- Type-safe state management
- Extensible node system
- Built-in retry logic
- Context support for cancellation
- Thread-safe execution

## Quick Start

```go
// Create a graph
g := graph.NewGraph("my-workflow")

// Add nodes
g.AddNode(modelNode)
g.AddNode(functionNode)

// Connect nodes
g.AddEdge(graph.NewEdge(modelNode.ID, functionNode.ID))

// Execute
result, err := g.Execute(ctx, initialState)
```

## Installation

```bash
go get hal
```
