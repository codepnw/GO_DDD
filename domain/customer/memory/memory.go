// a in-memory implementation of customer
package memory

import (
	"fmt"
	"sync"

	"github.com/codepnw/go-ddd/aggregate"
	"github.com/codepnw/go-ddd/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (m *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := m.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCusNotFound
}

func (m *MemoryRepository) Add(c aggregate.Customer) error {
	if m.customers == nil {
		m.Lock()
		m.customers = make(map[uuid.UUID]aggregate.Customer)
		m.Unlock()
	}
	// Make sure customer is already in repo
	if _, ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCus)
	}

	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}

func (m *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists: %w", customer.ErrUpdateCus)
	}

	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}