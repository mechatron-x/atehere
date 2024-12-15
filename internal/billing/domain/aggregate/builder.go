package aggregate

import (
	"errors"

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
			tableID:   uuid.Nil,
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

func (rcv *BillBuilder) SetTableID(tableID uuid.UUID) *BillBuilder {
	rcv.bill.tableID = tableID
	return rcv
}

func (rcv *BillBuilder) Build() (*Bill, error) {
	bill := rcv.bill

	if len(bill.billItems) == 0 {
		return nil, errors.New("bill items cannot be empty")
	}
	if bill.tableID == uuid.Nil {
		return nil, errors.New("table id cannot be nil")
	}

	return bill, nil
}
