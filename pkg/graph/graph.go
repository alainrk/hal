package graph

import (
	"fmt"
	"sync"
)

// Graph represents a directed graph of nodes
type Graph struct {
	mu          sync.RWMutex
	nodes       map[string]*Node
	edges       map[string][]*Edge
	entryPoint  string
	name        string
	description string
}

// NewGraph creates a new graph
func NewGraph(name string) *Graph {
	return &Graph{
		name:  name,
		nodes: make(map[string]*Node),
		edges: make(map[string][]*Edge),
	}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(node *Node) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.nodes[node.ID]; exists {
		return fmt.Errorf("node %s already exists", node.ID)
	}

	g.nodes[node.ID] = node

	// First node becomes entry point by default
	if g.entryPoint == "" {
		g.entryPoint = node.ID
	}

	return nil
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(edge *Edge) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.nodes[edge.From]; !exists {
		return fmt.Errorf("source node %s not found", edge.From)
	}
	if _, exists := g.nodes[edge.To]; !exists {
		return fmt.Errorf("target node %s not found", edge.To)
	}

	g.edges[edge.From] = append(g.edges[edge.From], edge)
	return nil
}

// SetEntryPoint sets the entry node for the graph
func (g *Graph) SetEntryPoint(nodeID string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, exists := g.nodes[nodeID]; !exists {
		return fmt.Errorf("node %s not found", nodeID)
	}

	g.entryPoint = nodeID
	return nil
}
