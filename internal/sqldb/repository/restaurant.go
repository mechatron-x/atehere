package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
)

type Restaurant struct {
	db     *gorm.DB
	mapper mapper.Restaurant
}

func NewRestaurant(db *gorm.DB) *Restaurant {
	return &Restaurant{
		db:     db,
		mapper: mapper.Restaurant{},
	}
}

func (r *Restaurant) Save(restaurant *aggregate.Restaurant) error {
	restaurantModel := r.mapper.FromAggregate(restaurant)

	result := r.db.Save(restaurantModel)

	return result.Error
}

func (r *Restaurant) GetByID(id uuid.UUID) (*aggregate.Restaurant, error) {
	var restaurantModel model.Restaurant

	result := r.db.
		Model(&model.Restaurant{}).
		Preload("WorkingDays").
		Preload("Tables").
		First(&restaurantModel, "id = ?", id.String())

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModel(&restaurantModel)
}

func (r *Restaurant) GetAll() ([]*aggregate.Restaurant, error) {
	restaurantModels := make([]model.Restaurant, 0)

	result := r.db.
		Model(&model.Restaurant{}).
		Preload("WorkingDays").
		Preload("Tables").
		Find(&restaurantModels)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModels(restaurantModels)
}

func (r *Restaurant) GetByOwnerID(ownerID uuid.UUID) ([]*aggregate.Restaurant, error) {
	models := make([]model.Restaurant, 0)

	result := r.db.
		Model(&model.Restaurant{}).
		Preload("WorkingDays").
		Preload("Tables").
		Find(&models, "owner_id = ?", ownerID.String())

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModels(models)
}
