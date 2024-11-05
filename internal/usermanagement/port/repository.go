package port

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type (
	CustomerRepository interface {
		Save(customer *aggregate.Customer) error
		GetByID(id uuid.UUID) (*aggregate.Customer, error)
	}

	ManagerRepository interface {
		Save(manager *aggregate.Manager) error
		GetByID(id uuid.UUID) (*aggregate.Manager, error)
	}
)
