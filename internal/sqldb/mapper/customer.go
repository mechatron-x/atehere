package mapper

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/sqldb/query"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type Customer struct{}

func NewUser() Customer {
	return Customer{}
}

func (c Customer) FromModel(model *query.Customer) (*aggregate.Customer, error) {
	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	customer := aggregate.NewCustomer()
	customer.SetID(id)
	customer.SetCreatedAt(model.CreatedAt.Time)
	customer.SetUpdatedAt(model.UpdatedAt.Time)

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	birthDate, err := valueobject.NewBirthDate(model.BirthDate.Time.Format(valueobject.BirthDateLayout))
	if err != nil {
		return nil, err
	}
	customer.SetFullName(fullName)
	customer.SetBirthDate(birthDate)

	return customer, nil
}

func (c Customer) FromAggregate(user *aggregate.Customer) *query.Customer {
	return &query.Customer{
		ID:       user.ID().String(),
		FullName: user.FullName().String(),
		BirthDate: sql.NullTime{
			Time:  user.BirthDate().Date(),
			Valid: true,
		},
		CreatedAt: sql.NullTime{
			Time:  user.CreatedAt(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  user.UpdatedAt(),
			Valid: true,
		},
		DeletedAt: sql.NullTime{
			Time:  user.DeletedAt(),
			Valid: user.IsDeleted(),
		},
	}
}
