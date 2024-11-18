package valueobject

import (
	"errors"
	"fmt"

	"github.com/mechatron-x/atehere/internal/core"
)

const (
	maxCategoryLength = 100
)

type Category struct {
	name string
}

func NewCategoryName(name string) (Category, error) {
	if core.IsEmptyString(name) {
		return Category{}, errors.New("category should not be empty")
	}

	if len(name) > maxCategoryLength {
		return Category{}, fmt.Errorf("category should not be longer than %d", maxCategoryLength)
	}

	return Category{name: name}, nil
}

func (c Category) String() string {
	return c.name
}
