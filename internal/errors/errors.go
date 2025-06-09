package errors

import (
	"fmt"
)

// GraphError represents a graph-related error
type GraphError struct {
	NodeID  string
	Message string
	Cause   error
}

// Error implements the error interface
func (e *GraphError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("graph error at node %s: %s: %v", e.NodeID, e.Message, e.Cause)
	}
	return fmt.Sprintf("graph error at node %s: %s", e.NodeID, e.Message)
}

// Unwrap returns the underlying error
func (e *GraphError) Unwrap() error {
	return e.Cause
}

// NewGraphError creates a new graph error
func NewGraphError(nodeID, message string, cause error) *GraphError {
	return &GraphError{
		NodeID:  nodeID,
		Message: message,
		Cause:   cause,
	}
}
