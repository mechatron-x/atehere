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

	BaseEvent struct {
		id         uuid.UUID
		invokeTime time.Time
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
