package port

import (
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type CustomerRepository interface {
	Save(customer *aggregate.Customer) error
	GetByID(id string) (*aggregate.Customer, error)
}

type ManagerRepository interface {
	Save(manager *aggregate.Manager) error
	GetByID(id string) (*aggregate.Manager, error)
}
