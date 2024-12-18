package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
)

type BillItemMapper struct{}

func NewBillItem() BillItemMapper {
	return BillItemMapper{}
}

func (bi BillItemMapper) FromModel(model *model.BillItem) (*entity.BillItem, error) {
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

	payments := make(map[uuid.UUID]valueobject.Price)
	for _, p := range model.Payments {
		verifiedCurrency, err := valueobject.ParseCurrency(p.Currency)
		if err != nil {
			return nil, err
		}

		verifiedPrice, err := valueobject.NewPrice(p.PaidPrice, verifiedCurrency)
		if err != nil {
			return nil, err
		}

		verifiedCustomerID, err := uuid.Parse(p.CustomerID)
		if err != nil {
			return nil, err
		}

		payments[verifiedCustomerID] = verifiedPrice
	}

	return entity.NewBillItemBuilder().
		SetID(model.ID).
		SetOwnerID(model.OwnerID).
		SetItemName(model.ItemName).
		SetUnitPrice(verifiedPrice).
		SetQuantity(verifiedQuantity).
		SetPayments(payments).
		SetCreatedAt(model.CreatedAt).
		SetUpdatedAt(model.UpdatedAt).
		Build()
}

func (bi BillItemMapper) FromModels(models []model.BillItem) ([]entity.BillItem, error) {
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

func (bi BillItemMapper) FromEntity(billID uuid.UUID, entity *entity.BillItem) *model.BillItem {
	payments := make([]model.BillItemPayments, 0)

	for customerID, price := range entity.Payments() {
		payment := model.BillItemPayments{
			CustomerID: customerID.String(),
			BillItemID: entity.ID().String(),
			PaidPrice:  price.Amount(),
			Currency:   price.Currency().String(),
		}

		payments = append(payments, payment)
	}

	return &model.BillItem{
		ID:        entity.ID().String(),
		OwnerID:   entity.OwnerID().String(),
		ItemName:  entity.ItemName(),
		UnitPrice: entity.UnitPrice().Amount(),
		Currency:  entity.UnitPrice().Currency().String(),
		Quantity:  entity.Quantity().Int(),
		Payments:  payments,
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
	}
}

func (bi BillItemMapper) FromEntities(billID uuid.UUID, entities []entity.BillItem) []model.BillItem {
	models := make([]model.BillItem, 0)
	for _, entity := range entities {
		billItem := bi.FromEntity(billID, &entity)
		models = append(models, *billItem)
	}
	return models
}
