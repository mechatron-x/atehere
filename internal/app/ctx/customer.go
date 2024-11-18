package ctx

import (
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
	"gorm.io/gorm"
)

type Customer struct {
	handler handler.Customer
}

func NewCustomer(db *gorm.DB, authenticator port.Authenticator) Customer {
	repo := repository.NewCustomer(db)
	service := service.NewCustomer(repo, authenticator)
	handler := handler.NewCustomerHandler(*service)

	return Customer{
		handler: handler,
	}
}

func (c Customer) Handler() handler.Customer {
	return c.handler
}
