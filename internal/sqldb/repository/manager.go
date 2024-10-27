package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

const (
	pkgManager = "repository.Manager"
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

func (m *Manager) Save(manager *aggregate.Manager) (*aggregate.Manager, core.PortError) {
	managerModel := m.mapper.FromAggregate(manager)
	saveParams := dal.SaveManagerParams(managerModel)

	managerModel, err := m.queries.SaveManager(context.Background(), saveParams)
	if err != nil {
		return nil, core.NewConnectionError(pkgManager, err)
	}

	return m.mapper.FromModel(managerModel)
}

func (m *Manager) GetByID(id string) (*aggregate.Manager, core.PortError) {
	managerModel, err := m.queries.GetManager(context.Background(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NewDataNotFoundError(pkgManager, err)
		}
		return nil, core.NewConnectionError(pkgManager, err)
	}

	return m.mapper.FromModel(managerModel)
}
