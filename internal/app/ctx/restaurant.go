package ctx

import (
	"database/sql"

	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/restaurant/port"
	"github.com/mechatron-x/atehere/internal/restaurant/service"
	"github.com/mechatron-x/atehere/internal/sqldb/repository"
)

type Restaurant struct {
	handler handler.Restaurant
}

func NewRestaurant(db *sql.DB, authenticator port.Authenticator, imageStorage port.ImageStorage, apiConf config.Api) Restaurant {
	repo := repository.NewRestaurant(db)
	service := service.NewRestaurant(repo, authenticator, imageStorage, apiConf)
	handler := handler.NewRestaurantHandler(*service)

	return Restaurant{
		handler: handler,
	}
}

func (r Restaurant) Handler() handler.Restaurant {
	return r.handler
}
