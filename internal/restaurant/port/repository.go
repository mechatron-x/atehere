package port

import (
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
)

type RestaurantRepository interface {
	Save(restaurant *aggregate.Restaurant) error
	GetByID(id string) (*aggregate.Restaurant, error)
	GetOwnerByID(id string) (*entity.Owner, error)
	GetAll(page, limit int) ([]*aggregate.Restaurant, error)
}
