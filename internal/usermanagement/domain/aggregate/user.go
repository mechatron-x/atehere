package aggregate

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type User struct {
	*core.Aggregate
	fullName  valueobject.FullName
	birthDate valueobject.BirthDate
}

func NewUser(fullName valueobject.FullName, birthDate valueobject.BirthDate) *User {
	return &User{
		Aggregate: core.DefaultAggregate(),
		fullName:  fullName,
		birthDate: birthDate,
	}
}
