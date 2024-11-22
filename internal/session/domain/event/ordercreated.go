package event

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
)

type (
	OrderCreated struct {
		core.BaseEvent
		tableID uuid.UUID
		orderID uuid.UUID
	}
)

func NewOrderCreated(tableID, orderID uuid.UUID) OrderCreated {
	return OrderCreated{
		BaseEvent: core.NewDomainEvent(),
		tableID:   tableID,
		orderID:   orderID,
	}
}

func (oc OrderCreated) TableID() uuid.UUID {
	return oc.tableID
}

func (oc OrderCreated) OrderID() uuid.UUID {
	return oc.orderID
}
