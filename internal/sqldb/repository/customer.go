package repository

import (
	"context"
	"database/sql"

	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type Customer struct {
	q *dal.Queries
	m mapper.Customer
}

func NewCustomer(db *sql.DB) *Customer {
	q := dal.New(db)
	m := mapper.NewUser()
	return &Customer{
		q: q,
		m: m,
	}
}

func (c *Customer) Save(user *aggregate.Customer) error {
	userModel := c.m.FromAggregate(user)
	saveParams := dal.SaveCustomerParams{
		ID:        userModel.ID,
		FullName:  userModel.FullName,
		BirthDate: userModel.BirthDate,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		DeletedAt: userModel.DeletedAt,
	}

	_, err := c.q.SaveCustomer(context.Background(), saveParams)
	return err
}

func (c *Customer) GetByID(id string) (*aggregate.Customer, error) {
	userModel, err := c.q.GetCustomer(context.Background(), id)
	if err != nil {
		return nil, err
	}

	user, err := c.m.FromModel(&userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}
