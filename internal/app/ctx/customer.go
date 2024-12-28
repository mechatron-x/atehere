package ctx

import (
	"github.com/mechatron-x/atehere/internal/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
	"gorm.io/gorm"
)

type Customer struct {
	handler handler.CustomerHandler
}

func NewCustomer(db *gorm.DB, authenticator port.Authenticator) Customer {
	repo := repository.NewCustomer(db)
	service := service.NewCustomer(repo, authenticator)
	handler := handler.NewCustomer(*service)

	return Customer{
		handler: handler,
	}
}

func (c Customer) Handler() handler.CustomerHandler {
	return c.handler
}
