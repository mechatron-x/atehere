package entity

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillItem struct {
	core.Entity
	owner      uuid.UUID
	itemName   string
	quantity   valueobject.Quantity
	unitPrice  valueobject.Price
	paidAmount valueobject.Price
}

func (rcv *BillItem) Owner() uuid.UUID {
	return rcv.owner
}

func (rcv *BillItem) ItemName() string {
	return rcv.itemName
}

func (rcv *BillItem) Quantity() valueobject.Quantity {
	return rcv.quantity
}

func (rcv *BillItem) UnitPrice() valueobject.Price {
	return rcv.unitPrice
}

func (rcv *BillItem) PaidAmount() valueobject.Price {
	return rcv.paidAmount
}

func (rcv *BillItem) RemainingAmount() valueobject.Price {
	return rcv.TotalDue().Subtract(rcv.paidAmount)
}

func (rcv *BillItem) TotalDue() valueobject.Price {
	return rcv.unitPrice.Multiply(float64(rcv.quantity))
}

func (rcv *BillItem) PayAll() {
	rcv.paidAmount = valueobject.NewPrice(0, rcv.unitPrice.Currency())
}
