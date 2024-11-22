package dto

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type (
	PlaceOrder struct {
		OrderedBy  string `json:"ordered_by"`
		MenuItemID string `json:"menu_item_id"`
		Quantity   int    `json:"quantity"`
	}
)

func (po *PlaceOrder) ToEntity() (entity.Order, error) {
	order := entity.NewOrder()

	verifiedMenuItemID, err := uuid.Parse(po.MenuItemID)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedQuantity, err := valueobject.NewQuantity(po.Quantity)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedOrderedBy, err := uuid.Parse(po.OrderedBy)
	if err != nil {
		return entity.Order{}, err
	}

	order.SetMenuItemID(verifiedMenuItemID)
	order.SetOrderedBy(verifiedOrderedBy)
	order.SetQuantity(verifiedQuantity)

	return order, nil
}
