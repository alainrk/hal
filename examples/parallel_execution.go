package main

import (
	"context"
	"fmt"
	"hal/pkg/graph"
	"hal/pkg/node"
	"hal/pkg/state"
	"log"
)

func parallelFlow() {
	// Example of parallel node execution
	g := graph.NewGraph("parallel-flow")

	// Create a fan-out node that triggers parallel execution
	fanOutNode := node.NewFunctionNode("fan-out", &node.FunctionNodeConfig{
		Function: func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
			return map[string]interface{}{
				"ready": true,
			}, nil
		},
		InputKeys: []string{},
	})

	// Create parallel processing nodes
	processor1 := createProcessor("processor-1", 1)
	processor2 := createProcessor("processor-2", 2)
	processor3 := createProcessor("processor-3", 3)

	// Create a fan-in node to collect results
	fanInNode := node.NewFunctionNode("fan-in", &node.FunctionNodeConfig{
		Function: func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
			var total int
			if v1, ok := input["result-1"]; ok {
				total += v1.(int)
			}
			if v2, ok := input["result-2"]; ok {
				total += v2.(int)
			}
			if v3, ok := input["result-3"]; ok {
				total += v3.(int)
			}

			return map[string]interface{}{
				"total": total,
			}, nil
		},
		InputKeys: []string{"result-1", "result-2", "result-3"},
	})

	// Add nodes
	g.AddNode(fanOutNode)
	g.AddNode(processor1)
	g.AddNode(processor2)
	g.AddNode(processor3)
	g.AddNode(fanInNode)

	// Connect for parallel execution
	g.AddEdge(graph.NewEdge(fanOutNode.ID, processor1.ID))
	g.AddEdge(graph.NewEdge(fanOutNode.ID, processor2.ID))
	g.AddEdge(graph.NewEdge(fanOutNode.ID, processor3.ID))

	// All processors connect to fan-in
	g.AddEdge(graph.NewEdge(processor1.ID, fanInNode.ID))
	g.AddEdge(graph.NewEdge(processor2.ID, fanInNode.ID))
	g.AddEdge(graph.NewEdge(processor3.ID, fanInNode.ID))

	// For true parallel execution, you would need to enhance the graph executor
	// This is a simplified example showing the structure

	initialState := state.NewState()
	ctx := context.Background()

	result, err := g.Execute(ctx, initialState)
	if err != nil {
		log.Fatal(err)
	}

	total, _ := result.State.Get("total")
	fmt.Printf("Total: %v\n", total)
}

func createProcessor(name string, value int) *graph.Node {
	return node.NewFunctionNode(name, &node.FunctionNodeConfig{
		Function: func(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
			// Simulate some work
			return map[string]interface{}{
				fmt.Sprintf("result-%d", value): value * 10,
			}, nil
		},
		InputKeys: []string{},
	})
}
