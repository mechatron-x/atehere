package core

import (
	"time"

	"github.com/google/uuid"
)

type Aggregate struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewAggregate() Aggregate {
	return Aggregate{
		id:        uuid.New(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: nil,
	}
}

func (a *Aggregate) ID() uuid.UUID {
	return a.id
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

func (a *Aggregate) IsDeleted() bool {
	return a.deletedAt != nil
}
