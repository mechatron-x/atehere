package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"gorm.io/gorm"
)

type Customer struct{}

func (c Customer) FromModel(model *model.Customer) (*aggregate.Customer, error) {
	customer := aggregate.NewCustomer()

	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	gender := valueobject.ParseGender(model.Gender)

	birthDate, err := valueobject.NewBirthDate(model.BirthDate)
	if err != nil {
		return nil, err
	}

	customer.SetID(id)
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

func (c Customer) FromAggregate(customer *aggregate.Customer) *model.Customer {
	return &model.Customer{
		ID:        customer.ID().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().Date().Format(valueobject.BirthDateLayout),
		Model: gorm.Model{
			CreatedAt: customer.CreatedAt(),
			UpdatedAt: customer.UpdatedAt(),
			DeletedAt: gorm.DeletedAt{
				Time:  customer.DeletedAt(),
				Valid: customer.IsDeleted(),
			},
		},
	}
}
