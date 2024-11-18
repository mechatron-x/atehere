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

	result := m.db.First(&model.Menu{ID: menuModel.ID})
	if result.Error == gorm.ErrRecordNotFound {
		result = m.db.Create(menuModel)
	} else {
		m.db.Model(menuModel).
			Association("MenuItems").
			Unscoped().
			Replace(menuModel.MenuItems)

		result = m.db.Updates(menuModel)
	}

	return result.Error
}

func (m *Menu) GetByID(id uuid.UUID) (*aggregate.Menu, error) {
	var menuModel model.Menu

	result := m.db.
		Model(&model.Menu{}).
		Preload("MenuItems").
		First(&menuModel, "id = ?", id.String())

	if result.Error != nil {
		return nil, result.Error
	}

	return m.mapper.FromModel(&menuModel)
}

func (m *Menu) GetManyByRestaurantID(restaurantID uuid.UUID) ([]*aggregate.Menu, error) {
	menuModels := make([]model.Menu, 0)

	result := m.db.
		Model(&model.Menu{}).
		Preload("MenuItems").
		Find(&menuModels, "restaurant_id = ?", restaurantID.String())

	if result.Error != nil {
		return nil, result.Error
	}

	return m.mapper.FromModels(menuModels)
}

func (m *Menu) IsRestaurantOwner(restaurantID, ownerID uuid.UUID) bool {
	restaurantModel := &model.Restaurant{
		ID:      restaurantID.String(),
		OwnerID: ownerID.String(),
	}

	result := m.db.First(restaurantModel)
	if result.Error != nil {
		return false
	}

	return result.RowsAffected != 0
}
