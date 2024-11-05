package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type Customer struct {
	queries *dal.Queries
	mapper  mapper.Customer
}

func NewCustomer(db *sql.DB) *Customer {
	return &Customer{
		queries: dal.New(db),
		mapper:  mapper.NewCustomer(),
	}
}

func (c *Customer) Save(customer *aggregate.Customer) error {
	customerModel := c.mapper.FromAggregate(customer)
	saveParams := dal.SaveCustomerParams(customerModel)

	err := c.queries.SaveCustomer(context.Background(), saveParams)
	if err != nil {
		return c.wrapError(err)
	}

	return nil
}

func (c *Customer) GetByID(id uuid.UUID) (*aggregate.Customer, error) {
	customerModel, err := c.queries.GetCustomer(context.Background(), id)
	if err != nil {
		return nil, c.wrapError(err)
	}

	return c.mapper.FromModel(customerModel)
}

func (c *Customer) wrapError(err error) error {
	return fmt.Errorf("repository.Customer: %v", err)
}
