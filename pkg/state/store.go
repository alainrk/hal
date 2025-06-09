package state

import (
	"fmt"
	"sync"
)

// Store provides persistent state storage
type Store interface {
	Save(id string, state *State) error
	Load(id string) (*State, error)
	Delete(id string) error
}

// MemoryStore is an in-memory state store
type MemoryStore struct {
	mu     sync.RWMutex
	states map[string]*State
}

// NewMemoryStore creates a new memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		states: make(map[string]*State),
	}
}

// Save stores a state
func (m *MemoryStore) Save(id string, state *State) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.states[id] = state.Clone()
	return nil
}

// Load retrieves a state
func (m *MemoryStore) Load(id string) (*State, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if state, exists := m.states[id]; exists {
		return state.Clone(), nil
	}

	return nil, fmt.Errorf("state %s not found", id)
}

// Delete removes a state
func (m *MemoryStore) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.states, id)
	return nil
}
