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
		MenuItemName  string  `json:"menu_item_name"`
		UnitPrice     float64 `json:"unit_price"`
		OrderPrice    float64 `json:"order_price"`
		PriceCurrency string  `json:"price_currency"`
		Quantity      int     `json:"quantity"`
	}

	SessionCustomer struct {
		ID       string `json:"customer_id"`
		FullName string `json:"full_name"`
	}

	TableOrders struct {
		Orders        []Order `json:"orders"`
		TotalPrice    float64 `json:"total_price"`
		PriceCurrency string  `json:"price_currency"`
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
