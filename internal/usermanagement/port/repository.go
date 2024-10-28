package port

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type CustomerRepository interface {
	Save(customer *aggregate.Customer) (*aggregate.Customer, core.PortError)
	GetByID(id string) (*aggregate.Customer, core.PortError)
}

type ManagerRepository interface {
	Save(manager *aggregate.Manager) (*aggregate.Manager, core.PortError)
	GetByID(id string) (*aggregate.Manager, core.PortError)
}
