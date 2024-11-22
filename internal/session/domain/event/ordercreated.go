package event

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
)

type (
	OrderCreated struct {
		core.BaseEvent
		sessionID uuid.UUID
		orderID   uuid.UUID
		quantity  int
	}
)

func NewOrderCreated(sessionID, orderID uuid.UUID, quantity int) OrderCreated {
	return OrderCreated{
		BaseEvent: core.NewDomainEvent(),
		sessionID: sessionID,
		orderID:   orderID,
		quantity:  quantity,
	}
}

func (oc OrderCreated) SessionID() uuid.UUID {
	return oc.sessionID
}

func (oc OrderCreated) OrderID() uuid.UUID {
	return oc.orderID
}

func (oc OrderCreated) Quantity() int {
	return oc.quantity
}
