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

func (rcv *Bill) Pay(paidBy, billItemID uuid.UUID, quantity valueobject.Quantity) error {
	for i, bi := range rcv.billItems {
		if bi.ID() != billItemID {
			continue
		}

		err := bi.Pay(paidBy, quantity)
		if err != nil {
			return err
		}

		rcv.billItems[i] = bi

		return nil
	}

	return fmt.Errorf("bill item with id %s not found", billItemID)
}
