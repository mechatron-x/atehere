package dto

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
)

type (
	MenuCreate struct {
		RestaurantID string `json:"restaurant_id"`
		Category     string `json:"category"`
	}

	MenuFilter struct {
		RestaurantID string `json:"restaurant_id"`
		Category     string `json:"category"`
		Name         string `json:"name"`
	}

	Menu struct {
		ID        string     `json:"id"`
		Category  string     `json:"category"`
		MenuItems []MenuItem `json:"menu_items"`
	}
)

func (mc MenuCreate) ToAggregate() (*aggregate.Menu, error) {
	verifiedCategory, err := valueobject.NewCategoryName(mc.Category)
	if err != nil {
		return nil, err
	}

	verifiedRestaurantID, err := uuid.Parse(mc.RestaurantID)
	if err != nil {
		return nil, err
	}

	menu := aggregate.NewMenu()
	menu.SetRestaurantID(verifiedRestaurantID)
	menu.SetCategory(verifiedCategory)

	return menu, nil
}

func ToMenu(menu *aggregate.Menu, imageCreator ImageURLCreatorFunc) *Menu {
	return &Menu{
		ID:        menu.ID().String(),
		Category:  menu.Category().String(),
		MenuItems: toMenuItemList(menu.MenuItems(), imageCreator),
	}
}

func ToMenuList(menus []*aggregate.Menu, imageCreator ImageURLCreatorFunc) []Menu {
	menuDtos := make([]Menu, 0)

	for _, m := range menus {
		menuDtos = append(menuDtos, *ToMenu(m, imageCreator))
	}

	return menuDtos
}
