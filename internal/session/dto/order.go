package dto

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type (
	PlaceOrders struct {
		TableID string        `json:"table_id"`
		Orders  []CreateOrder `json:"orders"`
	}

	CreateOrder struct {
		MenuItemID string `json:"menu_item_id"`
		Quantity   int    `json:"quantity"`
	}
)

func (po *PlaceOrders) ToEntities(orderedBy string) ([]entity.Order, error) {
	orders := make([]entity.Order, 0)

	for _, o := range po.Orders {
		order, err := o.ToEntity(orderedBy)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (co *CreateOrder) ToEntity(orderedBy string) (entity.Order, error) {
	order := entity.NewOrder()

	verifiedMenuItemID, err := uuid.Parse(co.MenuItemID)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedQuantity, err := valueobject.NewQuantity(co.Quantity)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedOrderedBy, err := uuid.Parse(orderedBy)
	if err != nil {
		return entity.Order{}, err
	}

	order.SetMenuItemID(verifiedMenuItemID)
	order.SetOrderedBy(verifiedOrderedBy)
	order.SetQuantity(verifiedQuantity)

	return order, nil
}
