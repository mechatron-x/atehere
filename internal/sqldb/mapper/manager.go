package mapper

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type Manager struct{}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) FromModel(model dal.Manager) (*aggregate.Manager, error) {
	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	manager := aggregate.NewManager()
	manager.SetID(id)
	manager.SetCreatedAt(model.CreatedAt.Time)
	manager.SetUpdatedAt(model.UpdatedAt.Time)

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	phoneNumber, err := valueobject.NewPhoneNumber(model.PhoneNumber)
	if err != nil {
		return nil, err
	}

	manager.SetFullName(fullName)
	manager.SetPhoneNumber(phoneNumber)

	return manager, nil
}

func (m Manager) FromAggregate(manager *aggregate.Manager) dal.Manager {
	return dal.Manager{
		ID:          manager.ID().String(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
		CreatedAt: sql.NullTime{
			Time:  manager.CreatedAt(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  manager.UpdatedAt(),
			Valid: true,
		},
		DeletedAt: sql.NullTime{
			Time:  manager.DeletedAt(),
			Valid: manager.IsDeleted(),
		},
	}
}
