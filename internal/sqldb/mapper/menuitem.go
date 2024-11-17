package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/menu/domain/entity"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
)

type MenuItem struct{}

func NewMenuItem() MenuItem {
	return MenuItem{}
}

func (mi MenuItem) FromModel(model *model.MenuItem) (*entity.MenuItem, error) {
	menuItem := entity.NewMenuItem()

	verifiedID, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	verifiedName, err := valueobject.NewMenuItemName(model.Name)
	if err != nil {
		return nil, err
	}

	verifiedImageName, err := valueobject.NewImage(model.ImageName)
	if err != nil {
		return nil, err
	}

	verifiedCurrency, err := valueobject.ParseCurrency(model.PriceCurrency)
	if err != nil {
		return nil, err
	}

	verifiedPrice := valueobject.NewPrice(model.PriceAmount, verifiedCurrency)

	verifiedDiscountPercentage, err := valueobject.NewPercentage(model.DiscountPercentage)
	if err != nil {
		return nil, err
	}

	menuItem.SetID(verifiedID)
	menuItem.SetName(verifiedName)
	menuItem.SetDescription(model.Description)
	menuItem.SetImageName(verifiedImageName)
	menuItem.SetPrice(verifiedPrice)
	menuItem.SetDiscountPercentage(verifiedDiscountPercentage)
	for _, i := range model.Ingredients {
		menuItem.AddIngredients(i.Ingredient)
	}

	return &menuItem, nil
}

func (mi MenuItem) FromModels(models []model.MenuItem) ([]*entity.MenuItem, error) {
	entities := make([]*entity.MenuItem, 0)
	for _, model := range models {
		entity, err := mi.FromModel(&model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (mi MenuItem) FromEntity(menuID uuid.UUID, entity *entity.MenuItem) *model.MenuItem {

	ingredients := make([]model.MenuItemIngredient, 0)
	for _, i := range entity.Ingredients() {
		menuItemIngredient := model.MenuItemIngredient{
			MenuItemID: entity.ID().String(),
			Ingredient: i,
		}

		ingredients = append(ingredients, menuItemIngredient)
	}

	return &model.MenuItem{
		ID:            entity.ID().String(),
		MenuID:        menuID.String(),
		Name:          entity.Name().String(),
		Description:   entity.Description(),
		ImageName:     entity.ImageName().String(),
		PriceAmount:   entity.Price().Amount(),
		PriceCurrency: entity.Price().Currency().String(),
		Ingredients:   ingredients,
		Model: gorm.Model{
			CreatedAt: entity.CreatedAt(),
			UpdatedAt: entity.UpdatedAt(),
		},
	}
}

func (mi MenuItem) FromEntities(menuID uuid.UUID, entities []*entity.MenuItem) []*model.MenuItem {
	models := make([]*model.MenuItem, 0)
	for _, entity := range entities {
		models = append(models, mi.FromEntity(menuID, entity))
	}

	return models
}
