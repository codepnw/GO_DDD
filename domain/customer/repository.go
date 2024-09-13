package customer

import (
	"errors"

	"github.com/codepnw/go-ddd/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCusNotFound    = errors.New("customer not found in the repository")
	ErrFailedToAddCus = errors.New("failed to add the customer")
	ErrUpdateCus      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
