package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mechatron-x/atehere/internal/core"
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
		return nil, core.ErrDbConnection
	}

	return c.mapper.FromModel(customerModel)
}

func (c *Customer) GetByID(id string) (*aggregate.Customer, error) {
	userModel, err := c.queries.GetCustomer(context.Background(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrModelNotFound
		}
		return nil, core.ErrDbConnection
	}

	return c.mapper.FromModel(userModel)
}
