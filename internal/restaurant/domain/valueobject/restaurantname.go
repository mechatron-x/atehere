package valueobject

import (
	"errors"
	"fmt"

	"github.com/mechatron-x/atehere/internal/core"
)

const (
	maxRestaurantNameLength = 100
)

type RestaurantName struct {
	name string
}

func NewRestaurantName(name string) (RestaurantName, error) {
	if core.IsEmptyString(name) {
		return RestaurantName{}, errors.New("restaurant name should not be empty")
	}

	if len(name) > maxRestaurantNameLength {
		return RestaurantName{}, fmt.Errorf("restaurant name should not be longer than %d", maxRestaurantNameLength)
	}

	return RestaurantName{name: name}, nil
}

func (rn RestaurantName) String() string {
	return rn.name
}
