package core

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
}

func NewEntity() Entity {
	return Entity{
		id:        uuid.New(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (e *Entity) ID() uuid.UUID {
	return e.id
}

func (e *Entity) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Entity) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Entity) SetID(id uuid.UUID) {
	e.id = id
}

func (e *Entity) SetCreatedAt(createdAt time.Time) {
	e.createdAt = createdAt
}

func (e *Entity) SetUpdatedAt(updatedAt time.Time) {
	e.updatedAt = updatedAt
}
