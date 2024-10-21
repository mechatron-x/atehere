package aggregate

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type User struct {
	core.Aggregate
	fullName  valueobject.FullName
	birthDate valueobject.BirthDate
}

func NewUser(fullName valueobject.FullName, birthDate valueobject.BirthDate) *User {
	return &User{
		Aggregate: core.NewAggregate(),
		fullName:  fullName,
		birthDate: birthDate,
	}
}

func LoadUser(root core.Aggregate, fullName valueobject.FullName, birthDate valueobject.BirthDate) *User {
	return &User{
		Aggregate: root,
		fullName:  fullName,
		birthDate: birthDate,
	}
}

func (u *User) FullName() valueobject.FullName {
	return u.fullName
}

func (u *User) BirthDate() valueobject.BirthDate {
	return u.birthDate
}
