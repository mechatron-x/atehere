package dto

import "github.com/mechatron-x/atehere/internal/restaurant/domain/entity"

type (
	TableCreate struct {
		Name string `json:"name"`
	}

	Table struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

func toTable(table entity.Table) Table {
	return Table{
		ID:   table.ID().String(),
		Name: table.Name().String(),
	}
}

func toTableList(tables []entity.Table) []Table {
	tableDtos := make([]Table, 0)
	for _, t := range tables {
		tableDtos = append(tableDtos, toTable(t))
	}

	return tableDtos
}
