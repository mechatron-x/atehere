package aggregate

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/menu/domain/entity"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
)

type Menu struct {
	core.Aggregate
	restaurantID uuid.UUID
	category     valueobject.Category
	menuItems    []entity.MenuItem
}

func NewMenu() *Menu {
	return &Menu{
		Aggregate: core.NewAggregate(),
		menuItems: make([]entity.MenuItem, 0),
	}
}

func (m *Menu) RestaurantID() uuid.UUID {
	return m.restaurantID
}

func (m *Menu) Category() valueobject.Category {
	return m.category
}

func (m *Menu) MenuItems() []entity.MenuItem {
	return m.menuItems
}

func (m *Menu) SetRestaurantID(restaurantID uuid.UUID) {
	m.restaurantID = restaurantID
}

func (m *Menu) SetCategory(category valueobject.Category) {
	m.category = category
}

func (m *Menu) AddMenuItems(menuItems ...entity.MenuItem) {
	for _, mi := range menuItems {
		m.addMenuItem(mi)
	}
}

func (m *Menu) DeleteMenuItem(id uuid.UUID) error {
	for i, mi := range m.menuItems {
		if mi.ID() == id {
			m.menuItems = slices.Delete(m.menuItems, i, i)
			return nil
		}
	}

	return fmt.Errorf("menu item with id: %s not found", id)
}

func (m *Menu) addMenuItem(menuItem entity.MenuItem) {
	for _, mi := range m.menuItems {
		if mi.ID() == menuItem.ID() {
			return
		}
	}

	m.menuItems = append(m.menuItems, menuItem)
}
