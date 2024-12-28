package ctx

import (
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
	"gorm.io/gorm"
)

type Manager struct {
	handler handler.ManagerHandler
}

func NewManager(db *gorm.DB, authenticator port.Authenticator) Manager {
	repo := repository.NewManager(db)
	service := service.NewManager(repo, authenticator)
	handler := handler.NewManager(*service)

	return Manager{
		handler: handler,
	}
}

func (m Manager) Handler() handler.ManagerHandler {
	return m.handler
}
