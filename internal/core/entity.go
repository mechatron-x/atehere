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

func LoadEntity(id uuid.UUID, createdAt, updatedAt time.Time) Entity {
	return Entity{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (e Entity) GetID() uuid.UUID {
	return e.id
}

func (e Entity) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e Entity) GetUpdatedAt() time.Time {
	return e.updatedAt
}
