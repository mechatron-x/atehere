package port

import (
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
)

type RestaurantRepository interface {
	Save(restaurant *aggregate.Restaurant) error
	// GetByID(id uuid.UUID) (*aggregate.Restaurant, error)
	// GetAll(page, limit int) ([]*aggregate.Restaurant, error)
}
