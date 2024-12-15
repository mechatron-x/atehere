package aggregate

import (
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

func (rcv *Bill) TotalDue(currency valueobject.Currency) valueobject.Price {
	totalDueAmount := valueobject.NewPrice(0, currency)

	for _, i := range rcv.billItems {
		totalDueAmount = totalDueAmount.Add(i.TotalDue())
	}

	return totalDueAmount
}
