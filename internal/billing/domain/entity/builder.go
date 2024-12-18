package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillItemBuilder struct {
	billItem *BillItem
	errs     []error
}

func NewBillItemBuilder() *BillItemBuilder {
	return &BillItemBuilder{
		billItem: &BillItem{
			Entity:   core.NewEntity(),
			payments: make(map[uuid.UUID]valueobject.Price),
		},
		errs: make([]error, 0),
	}
}

func (rcv *BillItemBuilder) SetID(id string) *BillItemBuilder {
	verifiedID, err := uuid.Parse(id)
	if err != nil {
		rcv.addError(err)
		return rcv
	}

	rcv.billItem.SetID(verifiedID)
	return rcv
}

func (rcv *BillItemBuilder) SetOwnerID(owner string) *BillItemBuilder {
	verifiedOwnerID, err := uuid.Parse(owner)
	if err != nil {
		rcv.addError(err)
		return rcv
	}

	rcv.billItem.ownerID = verifiedOwnerID
	return rcv
}

func (rcv *BillItemBuilder) SetItemName(itemName string) *BillItemBuilder {
	rcv.billItem.itemName = itemName
	return rcv
}

func (rcv *BillItemBuilder) SetQuantity(quantity valueobject.Quantity) *BillItemBuilder {
	rcv.billItem.quantity = quantity
	return rcv
}

func (rcv *BillItemBuilder) SetUnitPrice(price valueobject.Price) *BillItemBuilder {
	rcv.billItem.unitPrice = price
	return rcv
}

func (rcv *BillItemBuilder) SetPayments(payments map[uuid.UUID]valueobject.Price) *BillItemBuilder {
	rcv.billItem.payments = payments
	return rcv
}

func (rcv *BillItemBuilder) SetCreatedAt(createdAt time.Time) *BillItemBuilder {
	rcv.billItem.SetCreatedAt(createdAt)
	return rcv
}

func (rcv *BillItemBuilder) SetUpdatedAt(updatedAt time.Time) *BillItemBuilder {
	rcv.billItem.SetUpdatedAt(updatedAt)
	return rcv
}

func (rcv *BillItemBuilder) Build() (*BillItem, error) {
	if len(rcv.errs) != 0 {
		return nil, errors.Join(rcv.errs...)
	}

	return rcv.billItem, nil
}

func (rcv *BillItemBuilder) addError(err error) {
	rcv.errs = append(rcv.errs, err)
}
