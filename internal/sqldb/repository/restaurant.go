package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"gorm.io/gorm"
)

type Restaurant struct {
	db *gorm.DB
}

func NewRestaurant(db *gorm.DB) *Restaurant {
	return &Restaurant{
		db: db,
	}
}

func (r *Restaurant) Save(restaurant *aggregate.Restaurant) error {
	panic("not implemented")
}

func (r *Restaurant) GetByID(id uuid.UUID) (*aggregate.Restaurant, error) {
	panic("not implemented")
}

func (r *Restaurant) GetAll() ([]*aggregate.Restaurant, error) {
	panic("not implemented")
}

func (r *Restaurant) GetByOwnerID(ownerID uuid.UUID) ([]*aggregate.Restaurant, error) {
	panic("not implemented")
}
