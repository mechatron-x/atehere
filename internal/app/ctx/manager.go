package ctx

import (
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
	"gorm.io/gorm"
)

type Manager struct {
	handler handler.Manager
}

func NewManager(db *gorm.DB, authenticator port.Authenticator) Manager {
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
