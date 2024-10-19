package aggregate

import (
	"errors"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type User struct {
	*core.Aggregate
	fullName  valueobject.FullName
	birthDate valueobject.BirthDate
}

func NewUser(root *core.Aggregate, fullName valueobject.FullName, birthDate valueobject.BirthDate) (*User, error) {
	if root == nil {
		return nil, errors.New("aggregate root cannot be nil")
	}

	return &User{
		Aggregate: root,
		fullName:  fullName,
		birthDate: birthDate,
	}, nil
}
