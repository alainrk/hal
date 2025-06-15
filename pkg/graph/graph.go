package graph

import "log"

// Graph is the state machine that orchestrates the execution of Runnables.
type Graph[S any] struct {
	nodes      map[string]Runnable[S]
	edges      map[string]Router[S]
	entrypoint string
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

// SetEntryPoint defines the first node to be executed.
func (g *Graph[S]) SetEntryPoint(name string) {
	if g.entrypoint != "" {
		log.Fatalf("entrypoint is already set to '%s'", g.entrypoint)
	}
	g.entrypoint = name
}
