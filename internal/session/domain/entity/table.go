package entity

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
)

type Table struct {
	core.Entity
	restaurantID uuid.UUID
	name         string
}

func NewTable() Table {
	return Table{
		Entity: core.NewEntity(),
	}
}

func (t *Table) RestaurantID() uuid.UUID {
	return t.restaurantID
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) SetRestaurantID(restaurantID uuid.UUID) {
	t.restaurantID = restaurantID
}

func (t *Table) SetName(name string) {
	t.name = name
}
