package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type (
	Event struct {
		id         uuid.UUID
		eventType  valueobject.EventType
		invokeTime time.Time
	}
)

func NewEvent() Event {
	return Event{
		id:         uuid.New(),
		invokeTime: time.Now(),
	}
}

func (e *Event) ID() uuid.UUID {
	return e.id
}

func (e *Event) EventType() valueobject.EventType {
	return e.eventType
}

func (e *Event) InvokeTime() time.Time {
	return e.invokeTime
}

func (e *Event) SetID(id uuid.UUID) {
	e.id = id
}

func (e *Event) SetEventType(eventType valueobject.EventType) {
	e.eventType = eventType
}

func (e *Event) SetInvokeTime(invokeTime time.Time) {
	e.invokeTime = invokeTime
}
