package ctx

import (
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/menu/port"
	"github.com/mechatron-x/atehere/internal/menu/service"
	"gorm.io/gorm"
)

type Menu struct {
	handler handler.MenuHandler
}

func NewMenu(
	db *gorm.DB,
	authenticator port.Authenticator,
	imageStorage port.ImageStorage,
	apiConf config.Api,
) Menu {
	repo := repository.NewMenu(db)
	ms := service.NewMenu(repo, authenticator, imageStorage, apiConf)
	handler := handler.NewMenu(*ms)

	return Menu{
		handler: handler,
	}
}

func (m Menu) Handler() handler.MenuHandler {
	return m.handler
}
