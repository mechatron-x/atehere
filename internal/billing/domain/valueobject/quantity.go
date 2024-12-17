package valueobject

import (
	"errors"
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

func (q Quantity) Add(add Quantity) (Quantity, error) {
	result := q.Int() + add.Int()
	return NewQuantity(result)
}

func (q Quantity) Subtract(subtract Quantity) (Quantity, error) {
	remaining := q.Int() - subtract.Int()
	if remaining < 0 {
		return 0, errors.New("remaining quantity should not be smaller than 0")
	}

	return NewQuantity(remaining)
}

func (q Quantity) Equals(quantity Quantity) bool {
	return q.Int() == quantity.Int()
}

func (q Quantity) Compare(compare Quantity) int {
	if q.Int() == compare.Int() {
		return 0
	}

	if q.Int() > compare.Int() {
		return 1
	}

	return -1
}

func (q Quantity) Int() int {
	return int(q)
}
