package graph

// END is a special name for what it's the end node of every graph.
// When the router returns END, the graph execution stops.
const END = "__END__"

// Router determines the next node(s) to execute.
type Router[S any] interface {
	// Route inspects the current state and returns the name of the next node(s).
	// It can return langgraph.END to terminate the process, in which case END must be the only node returned.
	Route(state S) []string
}
