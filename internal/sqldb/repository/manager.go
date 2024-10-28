package repository

import (
	"context"
	"database/sql"

	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type Manager struct {
	queries *dal.Queries
	mapper  mapper.Manager
}

func NewManager(db *sql.DB) *Manager {
	return &Manager{
		queries: dal.New(db),
		mapper:  mapper.NewManager(),
	}
}

func (m *Manager) Save(manager *aggregate.Manager) error {
	managerModel := m.mapper.FromAggregate(manager)
	saveParams := dal.SaveManagerParams(managerModel)

	err := m.queries.SaveManager(context.Background(), saveParams)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetByID(id string) (*aggregate.Manager, error) {
	managerModel, err := m.queries.GetManager(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return m.mapper.FromModel(managerModel)
}
