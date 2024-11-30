package core

import (
	"sort"
	"time"

	"github.com/google/uuid"
)

type Aggregate struct {
	id        uuid.UUID
	events    []DomainEvent
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewAggregate() Aggregate {
	return Aggregate{
		id:        uuid.New(),
		events:    make([]DomainEvent, 0),
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: nil,
	}
}

func (a *Aggregate) ID() uuid.UUID {
	return a.id
}

func (a *Aggregate) Events() []DomainEvent {
	return a.events
}

func (a *Aggregate) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Aggregate) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Aggregate) DeletedAt() time.Time {
	if a.IsDeleted() {
		return *a.deletedAt
	}

	return time.Time{}
}

func (a *Aggregate) SetID(id uuid.UUID) {
	a.id = id
}

func (a *Aggregate) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func (a *Aggregate) SetUpdatedAt(updatedAt time.Time) {
	a.updatedAt = updatedAt
}

func (a *Aggregate) SetDeletedAt(deletedAt time.Time) {
	a.deletedAt = &deletedAt
}

func (a *Aggregate) RaiseEvent(event DomainEvent) {
	a.events = append(a.events, event)
	sort.SliceStable(a.events, func(i, j int) bool {
		return a.events[i].InvokeTime().Before(a.events[j].InvokeTime())
	})
}

func (a *Aggregate) ClearEvents() {
	a.events = make([]DomainEvent, 0)
}

func (a *Aggregate) IsDeleted() bool {
	return a.deletedAt != nil
}
