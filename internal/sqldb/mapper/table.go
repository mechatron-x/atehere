package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
)

type (
	Table struct{}
)

func NewTable() Table {
	return Table{}
}

func (rtm Table) FromModel(model dal.RestaurantTable) (entity.Table, error) {
	verifiedName, err := valueobject.NewTableName(model.Name)
	if err != nil {
		return entity.Table{}, err
	}

	table := entity.NewTable()
	table.SetID(model.ID)
	table.SetName(verifiedName)
	table.SetCreatedAt(model.CreatedAt)
	table.SetUpdatedAt(model.UpdatedAt)

	return table, nil
}

func (rtm Table) FromModels(models []dal.RestaurantTable) ([]entity.Table, error) {
	entities := make([]entity.Table, 0)
	for _, model := range models {
		entity, err := rtm.FromModel(model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (rtm Table) FromEntity(restaurantID uuid.UUID, entity entity.Table) dal.RestaurantTable {
	return dal.RestaurantTable{
		ID:           entity.ID(),
		RestaurantID: restaurantID,
		Name:         entity.Name().String(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
	}
}

func (rtm Table) FromEntities(restaurantID uuid.UUID, entities []entity.Table) []dal.RestaurantTable {
	models := make([]dal.RestaurantTable, 0)
	for _, entity := range entities {
		models = append(models, rtm.FromEntity(restaurantID, entity))
	}

	return models
}
