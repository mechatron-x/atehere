package dto

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	CheckoutEvent struct {
		ID           uuid.UUID
		RestaurantID string
		Table        string
		InvokeTime   int64
	}

	NewOrderEvent struct {
		ID           uuid.UUID
		RestaurantID string
		Table        string
		OrderedBy    string
		MenuItem     string
		Quantity     int
		InvokeTime   int64
	}
)

func (rcv CheckoutEvent) Message() string {
	return fmt.Sprintf("Table %s has requested checkout",
		rcv.Table,
	)
}

func (rcv NewOrderEvent) Message() string {
	return fmt.Sprintf("%s ordered %s x%d from table %s",
		rcv.OrderedBy,
		rcv.MenuItem,
		rcv.Quantity,
		rcv.Table,
	)
}
