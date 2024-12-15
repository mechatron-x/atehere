package aggregate

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/core"
)

type BillBuilder struct {
	errs []error
	bill *Bill
}

func NewBillBuilder() *BillBuilder {
	return &BillBuilder{
		errs: make([]error, 0),
		bill: &Bill{
			Aggregate: core.NewAggregate(),
			billItems: make([]entity.BillItem, 0),
			sessionID: uuid.Nil,
		},
	}
}

func (rcv *BillBuilder) SetID(id string) *BillBuilder {
	verifiedID, err := uuid.Parse(id)
	if err != nil {
		rcv.addError(err)
		return rcv
	}

	rcv.bill.SetID(verifiedID)
	return rcv
}

func (rcv *BillBuilder) SetBillItems(billItems []entity.BillItem) *BillBuilder {
	rcv.bill.billItems = billItems
	return rcv
}

func (rcv *BillBuilder) SetSessionID(sessionID string) *BillBuilder {
	verifiedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		rcv.addError(err)
		return rcv
	}
	rcv.bill.sessionID = verifiedSessionID
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

	if len(rcv.errs) != 0 {
		return nil, errors.Join(rcv.errs...)
	}

	if len(bill.billItems) == 0 {
		return nil, errors.New("bill items cannot be empty")
	}
	if bill.sessionID == uuid.Nil {
		return nil, errors.New("table id cannot be nil")
	}

	return bill, nil
}

func (rcv *BillBuilder) addError(err error) {
	rcv.errs = append(rcv.errs, err)
}
