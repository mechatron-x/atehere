package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"gorm.io/gorm"
)

type Customer struct {
	db     *gorm.DB
	mapper mapper.Customer
}

func NewCustomer(db *gorm.DB) *Customer {
	return &Customer{
		db:     db,
		mapper: mapper.Customer{},
	}
}

func (c *Customer) Save(customer *aggregate.Customer) error {
	model := c.mapper.FromAggregate(customer)

	result := c.db.Save(model)

	return result.Error
}

func (c *Customer) GetByID(id uuid.UUID) (*aggregate.Customer, error) {
	var model model.Customer

	result := c.db.First(&model, "id = ?", id.String())
	if result.Error != nil {
		return nil, result.Error
	}

	return c.mapper.FromModel(&model)
}
