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

	DomainEventHandler[TEvent DomainEvent] interface {
		HandleDomainEvent(event TEvent) error
	}
)

type (
	BaseEvent struct {
		id         uuid.UUID
		invokeTime time.Time
	}

	SessionClosedEvent struct {
		BaseEvent
		sessionID uuid.UUID
	}
)

func NewDomainEvent() BaseEvent {
	return BaseEvent{
		id:         uuid.New(),
		invokeTime: time.Now(),
	}
}

func (be BaseEvent) ID() uuid.UUID {
	return be.id
}

func (be BaseEvent) InvokeTime() time.Time {
	return be.invokeTime
}

func NewSessionClosedEvent(sessionID uuid.UUID) SessionClosedEvent {
	return SessionClosedEvent{
		BaseEvent: NewDomainEvent(),
		sessionID: sessionID,
	}
}

func (sc SessionClosedEvent) SessionID() uuid.UUID {
	return sc.sessionID
}
