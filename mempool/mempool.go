package mempool

import (
	"errors"
	"sync"

	"github.com/helixpay-xyz/go-helix/userop"
)

type Mempool struct {
	mu       sync.RWMutex
	all      []userop.UserOperation
	entities map[string][]userop.UserOperation
}

// NewMempool initializes a new Mempool.
func NewMempool() *Mempool {
	return &Mempool{
		all:      []userop.UserOperation{},
		entities: make(map[string][]userop.UserOperation),
	}
}

// AddOp adds multiple user operations to the mempool under a specific key.
func (m *Mempool) AddOp(key string, ops *[]userop.UserOperation) error {
	if ops == nil || len(*ops) == 0 {
		return errors.New("no operations to add")
	}

	m.mu.Lock() // Lock for writing
	defer m.mu.Unlock()

	if _, exists := m.entities[key]; !exists {
		m.entities[key] = []userop.UserOperation{}
	}

	for _, op := range *ops {
		m.all = append(m.all, op)
		m.entities[key] = append(m.entities[key], op)
	}

	return nil
}

// GetOps retrieves all user operations associated with a specific key.
func (m *Mempool) GetOps(key string) *[]userop.UserOperation {
	m.mu.RLock() // Lock for reading
	defer m.mu.RUnlock()

	ops, exists := m.entities[key]
	if !exists {
		return &[]userop.UserOperation{}
	}
	return &ops
}

// GetAllOps retrieves all user operations in the mempool.
func (m *Mempool) GetAllOps() *[]userop.UserOperation {
	m.mu.RLock() // Lock for reading
	defer m.mu.RUnlock()

	return &m.all
}

// Dump removes all operations from the mempool and returns them.
func (m *Mempool) Dump() *[]userop.UserOperation {
	m.mu.Lock() // Lock for writing
	defer m.mu.Unlock()

	dumpedOps := m.all
	m.all = []userop.UserOperation{}
	m.entities = make(map[string][]userop.UserOperation)
	return &dumpedOps
}
