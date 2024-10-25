package repository

import (
	"context"
	"database/sql"

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

func (c *Customer) Save(customer *aggregate.Customer) (*aggregate.Customer, error) {
	customerModel := c.mapper.FromAggregate(customer)
	saveParams := dal.SaveCustomerParams(customerModel)

	customerModel, err := c.queries.SaveCustomer(context.Background(), saveParams)
	if err != nil {
		return nil, err
	}

	return c.mapper.FromModel(customerModel)
}

func (c *Customer) GetByID(id string) (*aggregate.Customer, error) {
	userModel, err := c.queries.GetCustomer(context.Background(), id)
	if err != nil {
		return nil, err
	}

	user, err := c.mapper.FromModel(userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}
