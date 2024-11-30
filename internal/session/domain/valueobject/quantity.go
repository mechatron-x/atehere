package valueobject

import (
	"fmt"
)

type Quantity uint8

const maxQuantity = 250

func NewQuantity(quantity int) (Quantity, error) {
	if quantity < 0 {
		return 0, fmt.Errorf("quantity should not be smaller than 0")
	}
	if quantity > maxQuantity {
		return 0, fmt.Errorf("quantity should be smaller than %d", maxQuantity)
	}

	return Quantity(quantity), nil
}

func (q Quantity) Int() int {
	return int(q)
}
