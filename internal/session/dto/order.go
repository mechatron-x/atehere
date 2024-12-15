package dto

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type (
	CreateOrder struct {
		MenuItemID string `json:"menu_item_id"`
		Quantity   int    `json:"quantity"`
		OrderedBy  string `json:"ordered_by"`
	}

	PlaceOrders struct {
		Orders []CreateOrder `json:"orders"`
	}

	Order struct {
		CustomerFullName string  `json:"customer_full_name,omitempty"`
		MenuItemName     string  `json:"menu_item_name,omitempty"`
		Quantity         int     `json:"quantity,omitempty"`
		UnitPrice        float64 `json:"unit_price,omitempty"`
		TotalPrice       float64 `json:"total_price,omitempty"`
	}

	OrderList struct {
		Orders     []Order `json:"orders"`
		TotalPrice float64 `json:"total_price"`
		Currency   string  `json:"currency"`
	}
)

func (rcv *CreateOrder) ToEntity() (entity.Order, error) {
	order := entity.NewOrder()

	verifiedOrderedBy, err := uuid.Parse(rcv.OrderedBy)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedMenuItemID, err := uuid.Parse(rcv.MenuItemID)
	if err != nil {
		return entity.Order{}, err
	}

	verifiedQuantity, err := valueobject.NewQuantity(rcv.Quantity)
	if err != nil {
		return entity.Order{}, err
	}

	order.SetMenuItemID(verifiedMenuItemID)
	order.SetQuantity(verifiedQuantity)
	order.SetOrderedBy(verifiedOrderedBy)
	return order, nil
}

func (rcv *PlaceOrders) ToEntities(orderedBy string) ([]entity.Order, error) {
	orders := make([]entity.Order, 0)
	for _, o := range rcv.Orders {
		o.OrderedBy = orderedBy

		order, err := o.ToEntity()
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (rcv *OrderList) CalculateTotalPrice() {
	var totalPrice float64
	for _, o := range rcv.Orders {
		totalPrice += o.TotalPrice
	}

	rcv.TotalPrice = totalPrice
}
