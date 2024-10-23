package port

import "github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"

type CustomerRepository interface {
	Save(customer *aggregate.Customer) error
	GetByID(id string) (*aggregate.Customer, error)
}
