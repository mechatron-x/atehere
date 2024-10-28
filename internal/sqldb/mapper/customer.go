package mapper

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

const (
	pkgCustomer = "mapper.Customer"
)

type Customer struct{}

func NewCustomer() Customer {
	return Customer{}
}

func (c Customer) FromModel(model dal.Customer) (*aggregate.Customer, core.PortError) {
	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, core.NewDataMappingError(pkgCustomer, err)
	}

	customer := aggregate.NewCustomer()
	customer.SetID(id)
	customer.SetCreatedAt(model.CreatedAt.Time)
	customer.SetUpdatedAt(model.UpdatedAt.Time)

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, core.NewDataMappingError(pkgCustomer, err)
	}

	gender := valueobject.ParseGender(model.Gender)

	birthDate, err := valueobject.NewBirthDate(model.BirthDate.Time.Format(valueobject.BirthDateLayout))
	if err != nil {
		return nil, core.NewDataMappingError(pkgCustomer, err)
	}

	customer.SetFullName(fullName)
	customer.SetGender(gender)
	customer.SetBirthDate(birthDate)

	return customer, nil
}

func (c Customer) FromAggregate(user *aggregate.Customer) dal.Customer {
	return dal.Customer{
		ID:       user.ID().String(),
		FullName: user.FullName().String(),
		Gender:   user.Gender().String(),
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
