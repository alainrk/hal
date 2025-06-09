package main

import (
	"context"
	"fmt"
	"hal/pkg/graph"
	"hal/pkg/node"
	"hal/pkg/state"
	"log"
)

func conditionalFlow() {
	// Create a graph with conditional routing
	g := graph.NewGraph("conditional-flow")

	// Create nodes
	classifyNode := node.NewFunctionNode("classify", &node.FunctionNodeConfig{
		Function: func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
			text := input["text"].(string)

			// Simple classification logic
			category := "general"
			if len(text) > 100 {
				category = "long"
			} else {
				category = "short"
			}

			return map[string]interface{}{
				"category": category,
			}, nil
		},
		InputKeys: []string{"text"},
	})

	longHandler := node.NewFunctionNode("long-handler", &node.FunctionNodeConfig{
		Function: func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
			return map[string]interface{}{
				"result": "Processed long text",
			}, nil
		},
		InputKeys: []string{"text"},
	})

	shortHandler := node.NewFunctionNode("short-handler", &node.FunctionNodeConfig{
		Function: func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
			return map[string]interface{}{
				"result": "Processed short text",
			}, nil
		},
		InputKeys: []string{"text"},
	})

	// Add nodes
	g.AddNode(classifyNode)
	g.AddNode(longHandler)
	g.AddNode(shortHandler)

	// Add conditional edges
	g.AddEdge(graph.NewConditionalEdge(classifyNode.ID, longHandler.ID, func(s *state.State) bool {
		category, _ := s.Get("category")
		return category == "long"
	}))

	g.AddEdge(graph.NewConditionalEdge(classifyNode.ID, shortHandler.ID, func(s *state.State) bool {
		category, _ := s.Get("category")
		return category == "short"
	}))

	// Execute
	initialState := state.NewState()
	initialState.Set("text", "This is a test")

	ctx := context.Background()
	result, err := g.Execute(ctx, initialState)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v\n", result.State)
}
