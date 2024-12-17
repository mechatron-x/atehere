package entity

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillItem struct {
	core.Entity
	ownerID      uuid.UUID
	itemName     string
	unitPrice    valueobject.Price
	quantity     valueobject.Quantity
	paidQuantity valueobject.Quantity
	paidBy       uuid.UUIDs
}

func (rcv *BillItem) OwnerID() uuid.UUID {
	return rcv.ownerID
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

func (rcv *BillItem) PaidAmount() valueobject.Quantity {
	return rcv.paidQuantity
}

func (rcv *BillItem) PaidBy() uuid.UUIDs {
	return rcv.paidBy
}

func (rcv *BillItem) RemainingAmount() (valueobject.Quantity, error) {
	return rcv.quantity.Subtract(rcv.paidQuantity)
}

func (rcv *BillItem) TotalDue() (valueobject.Price, error) {
	return rcv.unitPrice.Multiply(float64(rcv.quantity))
}

func (rcv *BillItem) Pay(paidBy uuid.UUID, quantity valueobject.Quantity) error {
	pendingQuantity, err := rcv.paidQuantity.Add(quantity)
	if err != nil {
		return err
	}

	fmt.Println(pendingQuantity)

	if pendingQuantity.Compare(rcv.quantity) > 0 {
		return errors.New("requested payment quantity exceeds the remaining quantity")
	}

	rcv.paidBy = append(rcv.paidBy, paidBy)
	rcv.paidQuantity = pendingQuantity
	return nil
}

func (rcv *BillItem) IsAllPaid() bool {
	return rcv.quantity.Equals(rcv.paidQuantity)
}
