package aggregate

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/core"
)

type Bill struct {
	core.Aggregate
	tableID   uuid.UUID
	billItems []entity.BillItem
}

func (rcv *Bill) TableID() uuid.UUID {
	return rcv.tableID
}

func (rcv *Bill) BillItems() []entity.BillItem {
	return rcv.billItems
}
