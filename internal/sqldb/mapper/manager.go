package mapper

import (
	"database/sql"

	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type Manager struct{}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) FromModel(model dal.Manager) (*aggregate.Manager, error) {
	manager := aggregate.NewManager()

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	phoneNumber, err := valueobject.NewPhoneNumber(model.PhoneNumber)
	if err != nil {
		return nil, err
	}

	manager.SetID(model.ID)
	manager.SetFullName(fullName)
	manager.SetPhoneNumber(phoneNumber)
	manager.SetCreatedAt(model.CreatedAt)
	manager.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		manager.SetDeletedAt(model.DeletedAt.Time)
	}

	return manager, nil
}

func (m Manager) FromAggregate(manager *aggregate.Manager) dal.Manager {
	return dal.Manager{
		ID:          manager.ID(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
		CreatedAt:   manager.CreatedAt(),
		UpdatedAt:   manager.UpdatedAt(),
		DeletedAt: sql.NullTime{
			Time:  manager.DeletedAt(),
			Valid: manager.IsDeleted(),
		},
	}
}
