package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db     *gorm.DB
	mapper mapper.Customer
}

func NewCustomer(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db:     db,
		mapper: mapper.Customer{},
	}
}

func (c *CustomerRepository) Save(customer *aggregate.Customer) error {
	model := c.mapper.FromAggregate(customer)

	result := c.db.Save(model)

	return result.Error
}

func (c *CustomerRepository) GetByID(id uuid.UUID) (*aggregate.Customer, error) {
	customerModel := new(model.Customer)

	result := c.db.
		Where(&model.Customer{ID: id.String()}).
		First(customerModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return c.mapper.FromModel(customerModel)
}
