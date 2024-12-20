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
	domainEvent struct {
		id         uuid.UUID
		invokeTime time.Time
	}

	CheckoutEvent struct {
		domainEvent
		sessionID uuid.UUID
	}

	NewOrderEvent struct {
		domainEvent
		sessionID uuid.UUID
		orderID   uuid.UUID
		quantity  int
	}

	AllPaymentsDoneEvent struct {
		domainEvent
		sessionID uuid.UUID
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

func NewCheckoutEvent(sessionID uuid.UUID) CheckoutEvent {
	return CheckoutEvent{
		domainEvent: newDomainEvent(),
		sessionID:   sessionID,
	}
}

func (rcv CheckoutEvent) SessionID() uuid.UUID {
	return rcv.sessionID
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

func NewAllPaymentsDoneEvent(sessionId uuid.UUID) AllPaymentsDoneEvent {
	return AllPaymentsDoneEvent{
		sessionID: sessionId,
	}
}

func (rcv AllPaymentsDoneEvent) SessionID() uuid.UUID {
	return rcv.sessionID
}
