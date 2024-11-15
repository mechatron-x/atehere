package port

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/aggregate"
)

type MenuRepository interface {
	Save(menu *aggregate.Menu) error
	GetByID(id uuid.UUID) (*aggregate.Menu, error)
	IsRestaurantOwner(restaurantID, ownerID uuid.UUID) bool
	GetByCategory(category string) (*aggregate.Menu, error)
}
