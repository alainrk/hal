package graph

import "context"

type StateMerger[S any] interface {
	// Variadic states merger, with at least two states.
	Merge(context context.Context, states ...S) (S, error)
}
