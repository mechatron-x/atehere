package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"gorm.io/gorm"
)

type Manager struct{}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) FromModel(model *model.Manager) (*aggregate.Manager, error) {
	manager := aggregate.NewManager()

	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	phoneNumber, err := valueobject.NewPhoneNumber(model.PhoneNumber)
	if err != nil {
		return nil, err
	}

	manager.SetID(id)
	manager.SetFullName(fullName)
	manager.SetPhoneNumber(phoneNumber)
	manager.SetCreatedAt(model.CreatedAt)
	manager.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		manager.SetDeletedAt(model.DeletedAt.Time)
	}

	return manager, nil
}

func (m Manager) FromAggregate(manager *aggregate.Manager) *model.Manager {
	return &model.Manager{
		ID:          manager.ID().String(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
		Model: gorm.Model{
			CreatedAt: manager.CreatedAt(),
			UpdatedAt: manager.UpdatedAt(),
			DeletedAt: gorm.DeletedAt{
				Time:  manager.DeletedAt(),
				Valid: manager.IsDeleted(),
			},
		},
	}
}
