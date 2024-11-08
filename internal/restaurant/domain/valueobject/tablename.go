package valueobject

import (
	"errors"
	"fmt"

	"github.com/mechatron-x/atehere/internal/core"
)

const (
	maxTableNameLength int = 100
)

type TableName struct {
	name string
}

func NewTableName(name string) (TableName, error) {
	if core.IsEmptyString(name) {
		return TableName{}, errors.New("table name should not be empty")
	}

	if len(name) > maxTableNameLength {
		return TableName{}, fmt.Errorf("table name should not be longer than %d", maxTableNameLength)
	}

	return TableName{name: name}, nil
}

func (tn TableName) String() string {
	return tn.name
}
