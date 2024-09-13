package memory

import (
	"errors"
	"testing"

	"github.com/codepnw/go-ddd/aggregate"
	"github.com/codepnw/go-ddd/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name string
		id uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("john")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{id: cust},
	}

	testCases := []testCase{
		{
			name: "no customer by id",
			id: uuid.MustParse("fb4821f0-d6de-44bc-9b87-81d53c478d9b"),
			expectedErr: customer.ErrCusNotFound,
		}, {
			name: "customer by id",
			id: id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, get %v", tc.expectedErr, err)
			}
		})
	}
}
