package valueobject

import (
	"errors"
	"fmt"
)

type Percentage uint8

func NewPercentage(discountPercent int) (Percentage, error) {
	if discountPercent < 0 || discountPercent > 100 {
		return 0, errors.New("discount percent must be between 0 and 100")
	}

	return Percentage(discountPercent), nil
}

func (p Percentage) Amount() int {
	return int(p)
}

func (p Percentage) String() string {
	return fmt.Sprintf("%%	%d", p)
}
