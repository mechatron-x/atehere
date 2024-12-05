package dto

import (
	"fmt"

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

	OrderCreatedEvent struct {
		ID           uuid.UUID
		RestaurantID string
		Table        string
		OrderedBy    string
		MenuItem     string
		Quantity     int
		InvokeTime   int64
	}

	Order struct {
		MenuItemName string  `json:"menu_item_name"`
		UnitPrice    float64 `json:"unit_price"`
		OrderPrice   float64 `json:"order_price"`
		Currency     string  `json:"currency"`
		Quantity     int     `json:"quantity"`
	}

	CustomerOrders struct {
		OrderedBy      string  `json:"ordered_by"`
		CustomerOrders []Order `json:"customer_orders"`
	}

	OrdersView[TOrder OrderPriceCalculator] struct {
		Orders     []TOrder `json:"orders"`
		TotalPrice float64  `json:"total_price"`
		Currency   string   `json:"currency"`
	}

	SessionCustomer struct {
		ID       string `json:"customer_id"`
		FullName string `json:"full_name"`
	}
)

type (
	OrderPriceCalculator interface {
		GetOrderPrice() float64
		GetCurrency() string
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

func (rcv OrderCreatedEvent) Message() string {
	return fmt.Sprintf("%s ordered %s x%d from table %s",
		rcv.OrderedBy,
		rcv.MenuItem,
		rcv.Quantity,
		rcv.Table,
	)
}

func (rcv Order) GetOrderPrice() float64 {
	return rcv.OrderPrice
}

func (rcv Order) GetCurrency() string {
	return rcv.Currency
}

func (rcv CustomerOrders) GetOrderPrice() float64 {
	var orderPrice float64
	for _, o := range rcv.CustomerOrders {
		orderPrice += o.GetOrderPrice()
	}

	return orderPrice
}

func (rcv CustomerOrders) GetCurrency() string {
	if len(rcv.CustomerOrders) == 0 {
		return "N/A"
	}

	return rcv.CustomerOrders[0].GetCurrency()
}

func (rcv *OrdersView[TOrder]) CalculateTotalPrice() {
	var totalPrice float64
	var currency string
	for _, o := range rcv.Orders {
		totalPrice += o.GetOrderPrice()
		currency = o.GetCurrency()
	}

	rcv.TotalPrice = totalPrice
	rcv.Currency = currency
}
