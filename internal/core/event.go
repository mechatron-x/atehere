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
		menuItemID uuid.UUID
		orderedBy  uuid.UUID
		quantity   int
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

func NewOrder(menuItemID, orderedBy uuid.UUID, quantity int) Order {
	return Order{
		menuItemID: menuItemID,
		orderedBy:  orderedBy,
		quantity:   quantity,
	}
}

func (rcv Order) MenuItemID() uuid.UUID {
	return rcv.menuItemID
}

func (rcv Order) OrderedBy() uuid.UUID {
	return rcv.orderedBy
}

func (rcv Order) Quantity() int {
	return rcv.quantity
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
