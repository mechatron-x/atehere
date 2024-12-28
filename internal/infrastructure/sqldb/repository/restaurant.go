package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	db     *gorm.DB
	mapper mapper.Restaurant
}

func NewRestaurant(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{
		db:     db,
		mapper: mapper.Restaurant{},
	}
}

func (r *RestaurantRepository) Save(restaurant *aggregate.Restaurant) error {
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

	err = tx.Model(restaurantModel).
		Association("Locations").
		Unscoped().
		Replace(restaurantModel.Locations)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(restaurantModel)
	if result.Error != nil {
		tx.Rollback()
	}

	return result.Error
}

func (r *RestaurantRepository) GetByID(id uuid.UUID) (*aggregate.Restaurant, error) {
	restaurantModel := new(model.Restaurant)

	result := r.db.
		Preload("Tables").
		Preload("Locations").
		Where(&model.Restaurant{ID: id.String()}).
		First(restaurantModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModel(restaurantModel)
}

func (r *RestaurantRepository) GetAll() ([]*aggregate.Restaurant, error) {
	restaurantModels := make([]model.Restaurant, 0)

	result := r.db.
		Preload("Tables").
		Preload("Locations").
		Find(&restaurantModels)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModels(restaurantModels)
}

func (r *RestaurantRepository) GetByOwnerID(ownerID uuid.UUID) ([]*aggregate.Restaurant, error) {
	models := make([]model.Restaurant, 0)

	result := r.db.
		Preload("Tables").
		Preload("Locations").
		Where(&model.Restaurant{OwnerID: ownerID.String()}).
		Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	return r.mapper.FromModels(models)
}
