package aggregate

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillBuilder struct {
	bill *Bill
}

func NewBillBuilder() *BillBuilder {
	return &BillBuilder{
		bill: &Bill{
			Aggregate: core.NewAggregate(),
			billItems: make([]entity.BillItem, 0),
			sessionID: uuid.Nil,
		},
	}
}

func (rcv *BillBuilder) SetID(id uuid.UUID) *BillBuilder {
	rcv.bill.SetID(id)
	return rcv
}

func (rcv *BillBuilder) SetBillItems(billItems []entity.BillItem) *BillBuilder {
	rcv.bill.billItems = billItems
	return rcv
}

func (rcv *BillBuilder) SetSessionID(sessionID uuid.UUID) *BillBuilder {
	rcv.bill.sessionID = sessionID
	return rcv
}

func (rcv *BillBuilder) SetCreatedAt(createdAt time.Time) {
	rcv.bill.SetCreatedAt(createdAt)
}

func (rcv *BillBuilder) SetUpdatedAt(updatedAt time.Time) {
	rcv.bill.SetUpdatedAt(updatedAt)
}

func (rcv *BillBuilder) SetDeletedAt(deletedAt time.Time) {
	rcv.bill.SetDeletedAt(deletedAt)
}

func (rcv *BillBuilder) Build() (*Bill, error) {
	bill := rcv.bill

	if len(bill.billItems) == 0 {
		return nil, errors.New("bill items cannot be empty")
	}
	if bill.sessionID == uuid.Nil {
		return nil, errors.New("table id cannot be nil")
	}

	return bill, nil
}
