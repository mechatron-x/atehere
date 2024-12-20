package mapper

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/menu/domain/entity"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
)

type MenuItem struct{}

func NewMenuItem() MenuItem {
	return MenuItem{}
}

func (mi MenuItem) fromModel(model *model.MenuItem) (*entity.MenuItem, error) {
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
	menuItem.AddIngredients(model.Ingredients...)
	menuItem.SetCreatedAt(model.CreatedAt)
	menuItem.SetUpdatedAt(model.UpdatedAt)

	return &menuItem, nil
}

func (mi MenuItem) fromModels(models []model.MenuItem) ([]*entity.MenuItem, error) {
	entities := make([]*entity.MenuItem, 0)
	for _, model := range models {
		entity, err := mi.fromModel(&model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (mi MenuItem) fromEntity(menuID uuid.UUID, entity *entity.MenuItem) *model.MenuItem {
	return &model.MenuItem{
		ID:            entity.ID().String(),
		MenuID:        menuID.String(),
		Name:          entity.Name().String(),
		Description:   entity.Description(),
		ImageName:     entity.ImageName().String(),
		PriceAmount:   entity.Price().Amount(),
		PriceCurrency: entity.Price().Currency().String(),
		Ingredients:   pq.StringArray(entity.Ingredients()),
		CreatedAt:     entity.CreatedAt(),
		UpdatedAt:     entity.UpdatedAt(),
	}
}

func (mi MenuItem) fromEntities(menuID uuid.UUID, entities []entity.MenuItem) []model.MenuItem {
	models := make([]model.MenuItem, 0)
	for _, entity := range entities {
		menuItem := mi.fromEntity(menuID, &entity)
		models = append(models, *menuItem)
	}

	return models
}
