package entity

import (
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillItem struct {
	core.Entity
	name     string
	quantity valueobject.Quantity
	price    valueobject.Price
}

func (rcv *BillItem) Name() string {
	return rcv.name
}

func (rcv *BillItem) Quantity() valueobject.Quantity {
	return rcv.quantity
}

func (rcv *BillItem) Price() valueobject.Price {
	return rcv.price
}
