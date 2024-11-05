package ctx

import (
	"database/sql"

	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type Manager struct {
	handler handler.Manager
}

func NewManager(db *sql.DB, authenticator port.Authenticator) Manager {
	repo := repository.NewManager(db)
	service := service.NewManager(repo, authenticator)
	handler := handler.NewManagerHandler(*service)

	return Manager{
		handler: handler,
	}
}

func (m Manager) Handler() handler.Manager {
	return m.handler
}
