package graph

import (
	"context"
	"fmt"
	"hal/internal/utils"
	"log"

	"golang.org/x/sync/errgroup"
)

// Graph is the state machine that orchestrates the execution of Runnables.
type Graph[S any] struct {
	nodes       map[string]Runnable[S]
	edges       map[string]Router[S]
	entrypoint  string
	stateMerger StateMerger[S]
}

// NewGraph creates an empty Graph.
func NewGraph[S any]() *Graph[S] {
	return &Graph[S]{
		nodes: make(map[string]Runnable[S]),
		edges: make(map[string]Router[S]),
	}
}

// AddEdge adds a logic-based edge from a source node.
// The provided router will be used to dynamically select the next nodes based on the state.
func (g *Graph[S]) AddEdge(sourceName string, router Router[S]) {
	g.edges[sourceName] = router
}

// AddNode adds a runnable unit of work to the graph.
func (g *Graph[S]) AddNode(name string, node Runnable[S]) {
	if _, ok := g.nodes[name]; ok {
		log.Fatalf("a node called '%s' is already in the graph", name)
	}
	g.nodes[name] = node
}

// SetStateMerger sets the state merger function.
func (f *Graph[S]) SetStateMerger(merger StateMerger[S]) {
	f.stateMerger = merger
}

// SetEntryPoint defines the first node to be executed.
func (g *Graph[S]) SetEntryPoint(name string) {
	if g.entrypoint != "" {
		log.Fatalf("entrypoint is already set to '%s'", g.entrypoint)
	}
	g.entrypoint = name
}

// Invoke starts the execution of the graph.
func (g *Graph[S]) Invoke(ctx context.Context, initialState S) (S, error) {
	if _, ok := g.nodes[g.entrypoint]; !ok {
		log.Fatalf("entrypoint node not found in graph")
	}

	currentState := initialState

	// Nodes to run in the current step.
	stepNodes := []string{g.entrypoint}

	for len(stepNodes) > 0 {
		// Use an errgroup to manage parallel execution for the current step.
		errGroup, groupCtx := errgroup.WithContext(ctx)

		// A channel to safely collect results from each goroutine.
		type result struct {
			newState  S
			nextNodes []string
		}
		resultsChan := make(chan result, len(stepNodes))

		// Run each node independently.
		for _, nodeName := range stepNodes {
			name := nodeName // Clousure capture.
			errGroup.Go(func() error {
				node, ok := g.nodes[name]
				if !ok {
					return fmt.Errorf("node '%s' not found", name)
				}

				// Each node starts with the same state from the previous step.
				newState, err := node.Run(groupCtx, currentState)
				if err != nil {
					return fmt.Errorf("error in node '%s': %w", name, err)
				}

				var nextNodes []string
				if router, ok := g.edges[name]; ok {
					nextNodes = router.Route(newState)
				}

				select {
				case resultsChan <- result{newState: newState, nextNodes: nextNodes}:
					return nil
				case <-groupCtx.Done():
					return groupCtx.Err() // Handle context cancellation
				}
			})
		}

		// Wait for all nodes in the step to finish. If any failed, return the error.
		if err := errGroup.Wait(); err != nil {
			return initialState, err
		}
		close(resultsChan)

		// --- MERGE RESULTS ---
		var statesToMerge []S
		var allNextNodes []string

		for res := range resultsChan {
			statesToMerge = append(statesToMerge, res.newState)
			allNextNodes = append(allNextNodes, res.nextNodes...)
		}

		// If no nodes ran or produced results, we are done.
		if len(statesToMerge) == 0 {
			break
		}

		// Merge the states using the provided strategy.
		if len(statesToMerge) == 1 {
			currentState = statesToMerge[0]
		} else {
			if g.stateMerger == nil {
				return initialState, fmt.Errorf(
					"multiple nodes (%d) finished, but no StateMerger is configured",
					len(statesToMerge),
				)
			}

			var err error
			currentState, err = g.stateMerger.Merge(ctx, statesToMerge...)
			if err != nil {
				return initialState, fmt.Errorf("state merging failed: %w", err)
			}
		}

		// Prepare for the next step.
		stepNodes = utils.Unique(allNextNodes)
	}

	return currentState, nil
}
