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

	tx := r.db.Begin()
	defer tx.Commit()

	result := tx.First(&model.Restaurant{ID: restaurantModel.ID})

	if result.RowsAffected == 0 {
		result = tx.Create(restaurantModel)

		if result.Error != nil {
			tx.Rollback()
		}

		return result.Error
	}

	err := tx.Model(restaurantModel).
		Association("Tables").
		Unscoped().
		Replace(restaurantModel.Tables)
	if err != nil {
		tx.Rollback()
		return err
	}

	result = tx.Updates(restaurantModel)
	if result.Error != nil {
		tx.Rollback()
	}

	return result.Error
}

func (r *Restaurant) GetByID(id uuid.UUID) (*aggregate.Restaurant, error) {
	var restaurantModel model.Restaurant

	result := r.db.
		Model(&model.Restaurant{
			ID: id.String(),
		}).
		Preload("Tables").
		First(&restaurantModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModel(&restaurantModel)
}

func (r *Restaurant) GetAll() ([]*aggregate.Restaurant, error) {
	restaurantModels := make([]model.Restaurant, 0)

	result := r.db.
		Model(&model.Restaurant{}).
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
		Model(&model.Restaurant{
			OwnerID: ownerID.String(),
		}).
		Preload("Tables").
		Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModels(models)
}
