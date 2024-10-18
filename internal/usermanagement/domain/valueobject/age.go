package valueobject

import (
	"errors"
	"fmt"
)

type Age uint8

const (
	maxAge = int(^uint8(0)) // 255
)

func NewAge(age int) (Age, error) {
	if age < 0 {
		return 0, errors.New("age should not be smaller than 0")
	}

	if age > maxAge {
		return 0, fmt.Errorf("age should not be greater than %d", maxAge)
	}

	return Age(age), nil
}
