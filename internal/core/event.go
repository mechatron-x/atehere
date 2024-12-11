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

	SessionClosedEvent struct {
		domainEvent
		sessionID uuid.UUID
		orders    []Order
	}

	OrderCreatedEvent struct {
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

func NewSessionClosedEvent(sessionID uuid.UUID, orders []Order) SessionClosedEvent {
	return SessionClosedEvent{
		domainEvent: newDomainEvent(),
		sessionID:   sessionID,
		orders:      orders,
	}
}

func (rcv SessionClosedEvent) SessionID() uuid.UUID {
	return rcv.sessionID
}

func (rcv SessionClosedEvent) Orders() []Order {
	return rcv.orders
}

func NewOrderCreatedEvent(sessionID, orderID uuid.UUID, quantity int) OrderCreatedEvent {
	return OrderCreatedEvent{
		domainEvent: newDomainEvent(),
		sessionID:   sessionID,
		orderID:     orderID,
		quantity:    quantity,
	}
}

func (rcv OrderCreatedEvent) SessionID() uuid.UUID {
	return rcv.sessionID
}

func (rcv OrderCreatedEvent) OrderID() uuid.UUID {
	return rcv.orderID
}

func (rcv OrderCreatedEvent) Quantity() int {
	return rcv.quantity
}
