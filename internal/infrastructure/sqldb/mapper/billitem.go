package mapper

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
)

type BillItem struct{}

func NewBillItem() BillItem {
	return BillItem{}
}

func (bi BillItem) FromModel(model *model.BillItem) (*entity.BillItem, error) {
	verifiedQuantity, err := valueobject.NewQuantity(model.Quantity)
	if err != nil {
		return nil, err
	}

	verifiedCurrency, err := valueobject.ParseCurrency(model.Currency)
	if err != nil {
		return nil, err
	}

	verifiedPrice, err := valueobject.NewPrice(model.UnitPrice, verifiedCurrency)
	if err != nil {
		return nil, err
	}

	verifiedPaidQuantity, err := valueobject.NewQuantity(model.PaidQuantity)
	if err != nil {
		return nil, err
	}

	verifiedPaiBy := make([]uuid.UUID, 0)
	for _, i := range model.PaidBy {
		id, err := uuid.Parse(i)
		if err != nil {
			return nil, err
		}
		verifiedPaiBy = append(verifiedPaiBy, id)
	}

	return entity.NewBillItemBuilder().
		SetID(model.ID).
		SetOwnerID(model.OwnerID).
		SetItemName(model.ItemName).
		SetPrice(verifiedPrice).
		SetQuantity(verifiedQuantity).
		SetPaidQuantity(verifiedPaidQuantity).
		SetPaidBy(verifiedPaiBy).
		SetCreatedAt(model.CreatedAt).
		SetUpdatedAt(model.UpdatedAt).
		Build()
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
		ID:           entity.ID().String(),
		OwnerID:      entity.OwnerID().String(),
		ItemName:     entity.ItemName(),
		UnitPrice:    entity.UnitPrice().Amount(),
		Currency:     entity.UnitPrice().Currency().String(),
		Quantity:     entity.Quantity().Int(),
		PaidQuantity: entity.PaidAmount().Int(),
		PaidBy:       pq.StringArray(entity.PaidBy().Strings()),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
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
