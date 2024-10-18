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

func DefaultAggregate() *Aggregate {
	return &Aggregate{
		id:        uuid.New(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: nil,
	}
}

func NewAggregate(id uuid.UUID, createdAt, updatedAt time.Time, deletedAt *time.Time) *Aggregate {
	return &Aggregate{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

func (a *Aggregate) GetID() uuid.UUID {
	return a.id
}

func (a *Aggregate) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *Aggregate) GetUpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Aggregate) GetDeletedAt() (bool, time.Time) {
	if a.isDeleted() {
		return true, *a.deletedAt
	}

	return false, time.Time{}
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

func (a *Aggregate) isDeleted() bool {
	return a.deletedAt != nil
}
