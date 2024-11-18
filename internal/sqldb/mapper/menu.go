package mapper

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
)

type Menu struct {
	mi MenuItem
}

func NewMenu() Menu {
	return Menu{
		mi: NewMenuItem(),
	}
}

func (m Menu) FromModel(model *model.Menu) (*aggregate.Menu, error) {
	menu := aggregate.NewMenu()

	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	restaurantID, err := uuid.Parse(model.RestaurantID)
	if err != nil {
		return nil, err
	}

	category, err := valueobject.NewCategoryName(model.Category)
	if err != nil {
		return nil, err
	}

	menuItems, err := m.mi.FromModels(model.MenuItems)
	if err != nil {
		return nil, err
	}

	menu.SetID(id)
	menu.SetRestaurantID(restaurantID)
	menu.SetCategory(category)
	for _, mi := range menuItems {
		menu.AddMenuItems(*mi)
	}
	menu.SetCreatedAt(model.CreatedAt)
	menu.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		menu.SetDeletedAt(model.DeletedAt.Time)
	}

	return menu, nil
}

func (m Menu) FromModels(models []model.Menu) ([]*aggregate.Menu, error) {
	aggregates := make([]*aggregate.Menu, 0)

	for _, menu := range models {
		aggregate, err := m.FromModel(&menu)
		if err != nil {
			return nil, err
		}

		aggregates = append(aggregates, aggregate)
	}

	return aggregates, nil
}

func (m Menu) FromAggregate(aggregate *aggregate.Menu) *model.Menu {
	menuItems := make([]model.MenuItem, 0)
	for _, mi := range aggregate.MenuItems() {
		menuItem := m.mi.FromEntity(aggregate.ID(), &mi)

		menuItems = append(menuItems, *menuItem)
	}

	fmt.Println(aggregate.RestaurantID())
	return &model.Menu{
		ID:           aggregate.ID().String(),
		RestaurantID: aggregate.RestaurantID().String(),
		Category:     aggregate.Category().String(),
		MenuItems:    menuItems,
		Model: gorm.Model{
			CreatedAt: aggregate.CreatedAt(),
			UpdatedAt: aggregate.UpdatedAt(),
			DeletedAt: gorm.DeletedAt{
				Time:  aggregate.DeletedAt(),
				Valid: aggregate.IsDeleted(),
			},
		},
	}
}

func (m Menu) FromAggregates(aggregates []*aggregate.Menu) []*model.Menu {
	models := make([]*model.Menu, 0)
	for _, a := range aggregates {
		models = append(models, m.FromAggregate(a))
	}

	return models
}
