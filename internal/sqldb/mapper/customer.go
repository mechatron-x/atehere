package mapper

import (
	"database/sql"

	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type Customer struct{}

func NewCustomer() Customer {
	return Customer{}
}

func (c Customer) FromModel(model dal.Customer) (*aggregate.Customer, error) {
	customer := aggregate.NewCustomer()

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	gender := valueobject.ParseGender(model.Gender)

	birthDate, err := valueobject.NewBirthDate(model.BirthDate.Time.Format(valueobject.BirthDateLayout))
	if err != nil {
		return nil, err
	}

	customer.SetID(model.ID)
	customer.SetFullName(fullName)
	customer.SetGender(gender)
	customer.SetBirthDate(birthDate)
	customer.SetCreatedAt(model.CreatedAt)
	customer.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		customer.SetDeletedAt(model.DeletedAt.Time)
	}

	return customer, nil
}

func (c Customer) FromAggregate(customer *aggregate.Customer) dal.Customer {
	return dal.Customer{
		ID:       customer.ID(),
		FullName: customer.FullName().String(),
		Gender:   customer.Gender().String(),
		BirthDate: sql.NullTime{
			Time:  customer.BirthDate().Date(),
			Valid: true,
		},
		CreatedAt: customer.CreatedAt(),
		UpdatedAt: customer.UpdatedAt(),
		DeletedAt: sql.NullTime{
			Time:  customer.DeletedAt(),
			Valid: customer.IsDeleted(),
		},
	}
}
