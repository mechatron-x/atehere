package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
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

	verifiedCategory, err := valueobject.NewCategoryName(model.Category)
	if err != nil {
		return nil, err
	}

	verifiedMenuItems, err := m.mi.FromModels(model.MenuItems)
	if err != nil {
		return nil, err
	}

	menu.SetID(id)
	menu.SetCategory(verifiedCategory)
	for _, mi := range verifiedMenuItems {
		menu.AddMenuItems(*mi)
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

	return &model.Menu{
		ID:           aggregate.ID().String(),
		RestaurantID: aggregate.RestaurantID().String(),
		Category:     aggregate.Category().String(),
		MenuItems:    menuItems,
	}
}

func (m Menu) FromAggregates(aggregates []*aggregate.Menu) []*model.Menu {
	models := make([]*model.Menu, 0)
	for _, a := range aggregates {
		models = append(models, m.FromAggregate(a))
	}

	return models
}
