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

func LoadAggregate(id uuid.UUID, createdAt, updatedAt time.Time, deletedAt *time.Time) Aggregate {
	return Aggregate{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

func (a Aggregate) ID() uuid.UUID {
	return a.id
}

func (a Aggregate) CreatedAt() time.Time {
	return a.createdAt
}

func (a Aggregate) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a Aggregate) DeletedAt() time.Time {
	if a.IsDeleted() {
		return *a.deletedAt
	}

	return time.Time{}
}

func (a Aggregate) IsDeleted() bool {
	return a.deletedAt != nil
}
