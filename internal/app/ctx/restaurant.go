package ctx

import (
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/restaurant/port"
	"github.com/mechatron-x/atehere/internal/restaurant/service"
	"gorm.io/gorm"
)

type Restaurant struct {
	handler handler.RestaurantHandler
}

func NewRestaurant(
	db *gorm.DB,
	authenticator port.Authenticator,
	imageStorage port.ImageStorage,
	apiConf config.Api,
) Restaurant {
	repo := repository.NewRestaurant(db)
	service := service.NewRestaurant(repo, authenticator, imageStorage, apiConf)
	handler := handler.NewRestaurant(*service)

	return Restaurant{
		handler: handler,
	}
}

func (r Restaurant) Handler() handler.RestaurantHandler {
	return r.handler
}
