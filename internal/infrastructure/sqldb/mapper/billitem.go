package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
)

type BillItem struct{}

func NewBillItem() BillItem {
	return BillItem{}
}

func (bi BillItem) FromModel(model *model.BillItem) (*entity.BillItem, error) {
	builder := entity.NewBillItemBuilder()
	builder.SetID(model.ID)
	builder.SetOwnerID(model.OwnerID)

	verifiedQuantity, err := valueobject.NewQuantity(model.Quantity)
	if err != nil {
		return nil, err
	}
	builder.SetQuantity(verifiedQuantity)

	verifiedCurrency, err := valueobject.ParseCurrency(model.Currency)
	if err != nil {
		return nil, err
	}

	verifiedPrice := valueobject.NewPrice(model.UnitPrice, verifiedCurrency)
	builder.SetPrice(verifiedPrice)

	verifiedPaidPrice := valueobject.NewPrice(model.PaidPrice, verifiedCurrency)
	builder.SetPaidPrice(verifiedPaidPrice)

	builder.SetItemName(model.ItemName)
	builder.SetCreatedAt(model.CreatedAt)
	builder.SetUpdatedAt(model.UpdatedAt)

	return builder.Build()
}

func (bi BillItem) FromModels(models []model.BillItem) ([]entity.BillItem, error) {
	entities := make([]entity.BillItem, 0)
	for _, model := range models {
		entity, err := bi.FromModel(&model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}

func (bi BillItem) FromEntity(billID uuid.UUID, entity *entity.BillItem) *model.BillItem {
	return &model.BillItem{
		ID:        entity.ID().String(),
		OwnerID:   entity.OwnerID().String(),
		ItemName:  entity.ItemName(),
		Quantity:  entity.Quantity().Int(),
		UnitPrice: entity.UnitPrice().Amount(),
		PaidPrice: entity.PaidAmount().Amount(),
		Currency:  entity.UnitPrice().Currency().String(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
	}
}

func (bi BillItem) FromEntities(billID uuid.UUID, entities []entity.BillItem) []model.BillItem {
	models := make([]model.BillItem, 0)
	for _, entity := range entities {
		billItem := bi.FromEntity(billID, &entity)
		models = append(models, *billItem)
	}
	return models
}
