package valueobject

import (
	"errors"
	"fmt"

	"github.com/mechatron-x/atehere/internal/core"
)

const (
	maxMenuItemNameLength = 100
)

type MenuItemName struct {
	name string
}

func NewMenuItemName(name string) (MenuItemName, error) {
	if core.IsEmptyString(name) {
		return MenuItemName{}, errors.New("menu item name should not be empty")
	}

	if len(name) > maxMenuItemNameLength {
		return MenuItemName{}, fmt.Errorf("menu item name should not be longer than %d", maxMenuItemNameLength)
	}

	return MenuItemName{name: name}, nil
}

func (mn MenuItemName) String() string {
	return mn.name
}
