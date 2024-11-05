package port

import (
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
)

type RestaurantRepository interface {
	Save(restaurant *aggregate.Restaurant) error
	GetAll(page int) ([]*aggregate.Restaurant, error)
}
