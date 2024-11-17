package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
)

type Menu struct {
	db             *gorm.DB
	menuMapper     mapper.Menu
	menuItemMapper mapper.MenuItem
}

func NewMenu(db *gorm.DB) *Menu {
	return &Menu{
		db:             db,
		menuMapper:     mapper.Menu{},
		menuItemMapper: mapper.MenuItem{},
	}
}

func (m *Menu) Save(menu *aggregate.Menu) error {
	menuModel := m.menuMapper.FromAggregate(menu)

	result := m.db.Save(menuModel)

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

	return m.menuMapper.FromModel(&menuModel)
}

func (m *Menu) GetByRestaurantID(restaurantID uuid.UUID) ([]*aggregate.Menu, error) {
	menuModels := make([]model.Menu, 0)

	result := m.db.
		Model(&model.Menu{}).
		Preload("MenuItems").
		Find(&menuModels, "restaurant_id = ?", restaurantID.String())

	if result.Error != nil {
		return nil, result.Error
	}

	return m.menuMapper.FromModels(menuModels)
}

func (m *Menu) GetByCategory(restaurantID uuid.UUID, category string) (*aggregate.Menu, error) {
	var menuModel model.Menu

	result := m.db.
		Model(&model.Menu{}).
		Preload("MenuItems").
		First(&menuModel, "restaurant_id = ? AND category LIKE ?", restaurantID.String(), category)

	if result.Error != nil {
		return nil, result.Error
	}

	return m.menuMapper.FromModel(&menuModel)
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
