package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"gorm.io/gorm"
)

type ManagerRepository struct {
	db     *gorm.DB
	mapper mapper.Manager
}

func NewManager(db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{
		db:     db,
		mapper: mapper.NewManager(),
	}
}

func (m *ManagerRepository) Save(manager *aggregate.Manager) error {
	model := m.mapper.FromAggregate(manager)

	result := m.db.Save(model)

	return result.Error
}

func (m *ManagerRepository) GetByID(id uuid.UUID) (*aggregate.Manager, error) {
	managerModel := new(model.Manager)

	result := m.db.
		Where(&model.Manager{ID: id.String()}).
		First(managerModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return m.mapper.FromModel(managerModel)
}
