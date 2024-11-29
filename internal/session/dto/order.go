package dto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type (
	OrderCreate struct {
		MenuItemID string `json:"menu_item_id"`
		Quantity   int    `json:"quantity"`
	}

	OrderCreatedEventView struct {
		ID           uuid.UUID
		RestaurantID string
		Table        string
		OrderedBy    string
		MenuItem     string
		Quantity     int
		InvokeTime   int64
	}

	OrderCustomerView struct {
		MenuItemName string `json:"menu_item_name"`
		Quantity     int    `json:"quantity"`
	}

	OrderTableView struct {
		MenuItemName string `json:"menu_item_name"`
		Quantity     int    `json:"quantity"`
	}
)

func (po *OrderCreate) ToEntity() (entity.Order, error) {
	order := entity.NewOrder()

	verifiedMenuItemID, err := uuid.Parse(po.MenuItemID)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedQuantity, err := valueobject.NewQuantity(po.Quantity)
	if err != nil {
		return entity.Order{}, err
	}

	order.SetMenuItemID(verifiedMenuItemID)
	order.SetQuantity(verifiedQuantity)

	return order, nil
}

func (rcv OrderCreatedEventView) Message() string {
	return fmt.Sprintf("%s ordered %s x%d from table %s",
		rcv.OrderedBy,
		rcv.MenuItem,
		rcv.Quantity,
		rcv.Table,
	)
}
