package entity

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type Order struct {
	core.Entity
	menuItemID uuid.UUID
	orderedBy  uuid.UUID
	quantity   valueobject.Quantity
}

func NewOrder() Order {
	return Order{
		Entity: core.NewEntity(),
	}
}

func (o *Order) MenuItemID() uuid.UUID {
	return o.menuItemID
}

func (o *Order) OrderedBy() uuid.UUID {
	return o.orderedBy
}

func (o *Order) Quantity() valueobject.Quantity {
	return o.quantity
}

func (o *Order) SetMenuItemID(menuItemID uuid.UUID) {
	o.menuItemID = menuItemID
}

func (o *Order) SetOrderedBy(orderedBy uuid.UUID) {
	o.orderedBy = orderedBy
}

func (o *Order) SetQuantity(quantity valueobject.Quantity) {
	o.quantity = quantity
}

func (o *Order) AddQuantity(quantity valueobject.Quantity) error {
	newQuantity := o.quantity.Int() + quantity.Int()
	verifiedQuantity, err := valueobject.NewQuantity(newQuantity)
	if err != nil {
		return err
	}

	o.quantity = verifiedQuantity
	return nil
}
