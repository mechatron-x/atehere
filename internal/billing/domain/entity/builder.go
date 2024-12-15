package entity

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillItemBuilder struct {
	billItem BillItem
	errs     []error
}

func NewBillItemBuilder() *BillItemBuilder {
	return &BillItemBuilder{
		billItem: BillItem{
			Entity: core.NewEntity(),
		},
		errs: make([]error, 0),
	}
}

func (rcv *BillItemBuilder) SetID(id uuid.UUID) *BillItemBuilder {
	rcv.billItem.SetID(id)
	return rcv
}

func (rcv *BillItemBuilder) SetName(name string) *BillItemBuilder {
	rcv.billItem.name = name
	return rcv
}

func (rcv *BillItemBuilder) SetQuantity(quantity valueobject.Quantity) *BillItemBuilder {
	rcv.billItem.quantity = quantity
	return rcv
}

func (rcv *BillItemBuilder) SetPrice(price valueobject.Price) *BillItemBuilder {
	rcv.billItem.price = price
	return rcv
}

func (rcv *BillItemBuilder) Build() BillItem {
	return rcv.billItem
}
