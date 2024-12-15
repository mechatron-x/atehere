package core

import (
	"time"

	"github.com/google/uuid"
)

type (
	DomainEvent interface {
		ID() uuid.UUID
		InvokeTime() time.Time
	}
)

type (
	Order struct {
		MenuItemID uuid.UUID
		OrderedBy  uuid.UUID
		Quantity   int
	}
)

type (
	domainEvent struct {
		id         uuid.UUID
		invokeTime time.Time
	}

	CheckoutEvent struct {
		domainEvent
		sessionID uuid.UUID
		orders    []Order
	}

	NewOrderEvent struct {
		domainEvent
		sessionID uuid.UUID
		orderID   uuid.UUID
		quantity  int
	}
)

func newDomainEvent() domainEvent {
	return domainEvent{
		id:         uuid.New(),
		invokeTime: time.Now(),
	}
}

func (rcv domainEvent) ID() uuid.UUID {
	return rcv.id
}

func (rcv domainEvent) InvokeTime() time.Time {
	return rcv.invokeTime
}

func NewCheckoutEvent(sessionID uuid.UUID, orders []Order) CheckoutEvent {
	return CheckoutEvent{
		domainEvent: newDomainEvent(),
		sessionID:   sessionID,
		orders:      orders,
	}
}

func (rcv CheckoutEvent) SessionID() uuid.UUID {
	return rcv.sessionID
}

func (rcv CheckoutEvent) Orders() []Order {
	return rcv.orders
}

func NewOrderCreatedEvent(sessionID, orderID uuid.UUID, quantity int) NewOrderEvent {
	return NewOrderEvent{
		domainEvent: newDomainEvent(),
		sessionID:   sessionID,
		orderID:     orderID,
		quantity:    quantity,
	}
}

func (rcv NewOrderEvent) SessionID() uuid.UUID {
	return rcv.sessionID
}

func (rcv NewOrderEvent) OrderID() uuid.UUID {
	return rcv.orderID
}

func (rcv NewOrderEvent) Quantity() int {
	return rcv.quantity
}
