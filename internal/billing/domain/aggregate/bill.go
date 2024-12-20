package aggregate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type Bill struct {
	core.Aggregate
	sessionID uuid.UUID
	billItems []entity.BillItem
}

func (rcv *Bill) SessionID() uuid.UUID {
	return rcv.sessionID
}

func (rcv *Bill) BillItems() []entity.BillItem {
	return rcv.billItems
}

func (rcv *Bill) Pay(paidBy, billItemID uuid.UUID, price valueobject.Price) error {
	defer rcv.allPaymentsDonePolicy()

	for i, bi := range rcv.billItems {
		if bi.ID() != billItemID {
			continue
		}

		err := bi.Pay(paidBy, price)
		if err != nil {
			return err
		}

		rcv.billItems[i] = bi
		return nil
	}

	return fmt.Errorf("bill item with id %s not found", billItemID)
}

func (rcv *Bill) allPaymentsDonePolicy() {
	for _, bi := range rcv.billItems {
		if !bi.IsPaid() {
			return
		}
	}

	print("all orders paid")
	event := core.NewAllPaymentsDoneEvent(rcv.sessionID)
	rcv.RaiseEvent(event)
}
