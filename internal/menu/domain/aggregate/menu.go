package aggregate

import (
	"fmt"
	"slices"
	"strings"

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

func (m *Menu) SetMenuItems(menuItems ...entity.MenuItem) {
	m.menuItems = append(m.menuItems, menuItems...)
}

func (m *Menu) AddMenuItems(menuItems ...entity.MenuItem) error {
	newItems := make([]entity.MenuItem, 0)
	for _, mi := range menuItems {
		if err := m.addMenuItemPolicy(mi); err != nil {
			return err
		}

		newItems = append(newItems, mi)
	}

	m.menuItems = append(m.menuItems, newItems...)
	return nil
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

func (m *Menu) addMenuItemPolicy(menuItem entity.MenuItem) error {
	for _, mi := range m.menuItems {
		if mi.ID() == menuItem.ID() {
			return fmt.Errorf("menu item with id %s already exists", menuItem.ID())
		}

		if strings.Compare(mi.Name().String(), menuItem.Name().String()) == 0 {
			return fmt.Errorf("menu item with name %s already exists", menuItem.Name())
		}
	}

	return nil
}
