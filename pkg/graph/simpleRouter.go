package graph

// AddSimpleEdge creates a simple, static unconditional edge from a source node to a destination node.
func (g *Graph[S]) AddSimpleEdge(sourceName, destinationName string) {
	// We use a simple, internal router that always returns the same destination.
	router := simpleRouter[S]{destination: destinationName}
	g.edges[sourceName] = router
}

// simpleRouter is a private helper struct that implements the Router interface
// for unconditional edges created with AddEdge.
type simpleRouter[S any] struct {
	destination string
}

func (r simpleRouter[S]) Route(_ S) []string {
	return []string{r.destination}
}
