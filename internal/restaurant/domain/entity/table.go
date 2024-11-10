package entity

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
)

type Table struct {
	core.Entity
	name valueobject.TableName
}

func NewTable() Table {
	return Table{
		Entity: core.NewEntity(),
	}
}

func (t *Table) Name() valueobject.TableName {
	return t.name
}

func (t *Table) SetName(name valueobject.TableName) {
	t.name = name
}
