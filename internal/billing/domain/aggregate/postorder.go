package aggregate

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/core"
)

type PostOrder struct {
	core.Aggregate
	sessionID  uuid.UUID
	orderedBy  uuid.UUID
	menuItemID uuid.UUID
	quantity   valueobject.Quantity
}

func NewPostOrder() *PostOrder {
	return &PostOrder{
		Aggregate: core.NewAggregate(),
	}
}

func (rcv *PostOrder) SessionID() uuid.UUID {
	return rcv.sessionID
}

func (rcv *PostOrder) OrderedBy() uuid.UUID {
	return rcv.orderedBy
}

func (rcv *PostOrder) MenuItemID() uuid.UUID {
	return rcv.menuItemID
}

func (rcv *PostOrder) Quantity() valueobject.Quantity {
	return rcv.quantity
}

func (rcv *PostOrder) SetSessionID(sessionID uuid.UUID) {
	rcv.sessionID = sessionID
}

func (rcv *PostOrder) SetOrderedBy(orderedBy uuid.UUID) {
	rcv.orderedBy = orderedBy
}

func (rcv *PostOrder) SetMenuItemID(menuItemID uuid.UUID) {
	rcv.menuItemID = menuItemID
}

func (rcv *PostOrder) SetQuantity(quantity valueobject.Quantity) {
	rcv.quantity = quantity
}
