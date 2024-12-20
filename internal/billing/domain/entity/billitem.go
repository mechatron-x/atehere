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
	ownerID   uuid.UUID
	itemName  string
	unitPrice valueobject.Price
	quantity  valueobject.Quantity
	payments  map[uuid.UUID]valueobject.Price
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

func (rcv *BillItem) Payments() map[uuid.UUID]valueobject.Price {
	return rcv.payments
}

func (rcv *BillItem) TotalDue() valueobject.Price {
	return rcv.unitPrice.Multiply(float64(rcv.quantity))
}

func (rcv *BillItem) PaidAmount() valueobject.Price {
	paidTotal := valueobject.MustPrice(0, rcv.unitPrice.Currency())
	for _, price := range rcv.payments {
		paidTotal = paidTotal.Add(price)
	}

	return paidTotal
}

func (rcv *BillItem) RemainingAmount() valueobject.Price {
	remainingAmount, err := rcv.TotalDue().Subtract(rcv.PaidAmount())
	if err != nil {
		return valueobject.MustPrice(0, rcv.unitPrice.Currency())
	}

	return remainingAmount
}

func (rcv *BillItem) Pay(paidBy uuid.UUID, price valueobject.Price) error {
	if rcv.TotalDue().Equals(rcv.PaidAmount()) {
		return errors.New("payment failed: item already paid")
	}

	pendingPrice := rcv.PaidAmount().Add(price)
	_, err := rcv.TotalDue().Subtract(pendingPrice)
	if err != nil {
		return fmt.Errorf("payment failed: %v", err)
	}

	rcv.addCustomerPayments(paidBy, price)
	return nil
}

func (rcv *BillItem) addCustomerPayments(paidBy uuid.UUID, price valueobject.Price) {
	paidPrice, ok := rcv.payments[paidBy]
	if ok {
		rcv.payments[paidBy] = paidPrice.Add(price)
		return
	}

	rcv.payments[paidBy] = price
}

func (rcv *BillItem) IsPaid() bool {
	return rcv.PaidAmount().Equals(rcv.TotalDue())
}
