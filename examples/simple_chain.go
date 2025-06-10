package main

import (
	"context"
	"fmt"
	"hal/pkg/graph"
	"hal/pkg/model"
	"hal/pkg/node"
	"hal/pkg/state"
	"log"
)

func simpleChain() {
	// Create a new graph
	g := graph.NewGraph("simple-chain")

	// Create model (implement your preferred model)
	llm := model.NewOpenAIModel("your-api-key", "gpt-4", "")

	// Create nodes
	summarizeNode := node.NewModelNode("summarize", &node.ModelNodeConfig{
		Model:        llm,
		PromptKey:    "input_text",
		OutputKey:    "summary",
		SystemPrompt: "You are a helpful assistant that creates concise summaries.",
		Temperature:  0.7,
	})

	analyzeNode := node.NewModelNode("analyze", &node.ModelNodeConfig{
		Model:        llm,
		PromptKey:    "summary",
		OutputKey:    "analysis",
		SystemPrompt: "You are an expert analyst.",
		Temperature:  0.5,
	})

	// Add nodes to graph
	g.AddNode(summarizeNode)
	g.AddNode(analyzeNode)

	// Connect nodes
	g.AddEdge(graph.NewEdge(summarizeNode.ID, analyzeNode.ID))

	// Create initial state
	initialState := state.NewState()
	initialState.Set("input_text", "Your long text here...")

	// Execute graph
	ctx := context.Background()
	result, err := g.Execute(ctx, initialState)
	if err != nil {
		log.Fatal(err)
	}

	// Get results
	analysis, _ := result.State.Get("analysis")
	fmt.Printf("Analysis: %v\n", analysis)
}
