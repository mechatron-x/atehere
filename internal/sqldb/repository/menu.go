package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
)

type Menu struct {
	db     *gorm.DB
	mapper mapper.Menu
}

func NewMenu(db *gorm.DB) *Menu {
	return &Menu{
		db:     db,
		mapper: mapper.Menu{},
	}
}

func (m *Menu) Save(menu *aggregate.Menu) error {
	menuModel := m.mapper.FromAggregate(menu)

	tx := m.db.Begin()
	defer tx.Commit()

	result := tx.First(&model.Menu{ID: menuModel.ID})
	if result.RowsAffected == 0 {
		result = tx.Create(menuModel)

		if result.Error != nil {
			tx.Rollback()
		}

		return result.Error
	}

	err := tx.Model(menuModel).
		Association("MenuItems").
		Unscoped().
		Replace(menuModel.MenuItems)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(menuModel)
	if result.Error != nil {
		tx.Rollback()
	}

	return result.Error
}

func (m *Menu) GetByID(id uuid.UUID) (*aggregate.Menu, error) {
	menuModel := new(model.Menu)

	result := m.db.
		Preload("MenuItems").
		Where(&model.Menu{ID: id.String()}).
		First(menuModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return m.mapper.FromModel(menuModel)
}

func (m *Menu) GetManyByRestaurantID(restaurantID uuid.UUID) ([]*aggregate.Menu, error) {
	menuModels := make([]model.Menu, 0)

	result := m.db.
		Preload("MenuItems").
		Where(&model.Menu{RestaurantID: restaurantID.String()}).
		Find(&menuModels)

	if result.Error != nil {
		return nil, result.Error
	}

	return m.mapper.FromModels(menuModels)
}

func (m *Menu) IsRestaurantOwner(restaurantID, ownerID uuid.UUID) bool {
	restaurantModel := new(model.Restaurant)

	result := m.db.
		Where(&model.Restaurant{
			ID:      restaurantID.String(),
			OwnerID: ownerID.String(),
		}).
		First(restaurantModel)
	if result.Error != nil {
		return false
	}

	return result.RowsAffected != 0
}
