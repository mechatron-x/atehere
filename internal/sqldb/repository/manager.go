package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"gorm.io/gorm"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (m *Manager) Save(manager *aggregate.Manager) error {
	panic("not implemented")
}

func (m *Manager) GetByID(id uuid.UUID) (*aggregate.Manager, error) {
	panic("not implemented")
}
