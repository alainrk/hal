package state

import (
	"encoding/json"
	"sync"
)

// State represents the execution state
type State struct {
	mu      sync.RWMutex
	data    map[string]any
	history []StateSnapshot
}

// StateSnapshot represents a point-in-time state
type StateSnapshot struct {
	NodeID    string
	Timestamp int64
	Data      map[string]any
}

// NewState creates a new state
func NewState() *State {
	return &State{
		data:    make(map[string]any),
		history: make([]StateSnapshot, 0),
	}
}

// Get retrieves a value from state
func (s *State) Get(key string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, exists := s.data[key]
	return val, exists
}

// Set stores a value in state
func (s *State) Set(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

// Update merges new data into state
func (s *State) Update(updates map[string]any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for k, v := range updates {
		s.data[k] = v
	}
}

// Clone creates a deep copy of the state
func (s *State) Clone() *State {
	s.mu.RLock()
	defer s.mu.RUnlock()

	newState := NewState()

	// Deep copy via JSON marshaling
	data, _ := json.Marshal(s.data)
	json.Unmarshal(data, &newState.data)

	return newState
}
