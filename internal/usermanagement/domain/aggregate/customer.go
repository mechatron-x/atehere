package aggregate

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type Customer struct {
	core.Aggregate
	email     valueobject.Email
	password  valueobject.Password
	fullName  valueobject.FullName
	gender    valueobject.Gender
	birthDate valueobject.BirthDate
}

func NewCustomer() *Customer {
	return &Customer{
		Aggregate: core.NewAggregate(),
	}
}

func (c *Customer) Email() valueobject.Email {
	return c.email
}

func (c *Customer) FullName() valueobject.FullName {
	return c.fullName
}

func (c *Customer) Gender() valueobject.Gender {
	return c.gender
}

func (c *Customer) BirthDate() valueobject.BirthDate {
	return c.birthDate
}

func (c *Customer) SetEmail(email valueobject.Email) {
	c.email = email
}

func (c *Customer) SetPassword(password valueobject.Password) {
	c.password = password
}

func (c *Customer) SetFullName(fullName valueobject.FullName) {
	c.fullName = fullName
}

func (c *Customer) SetGender(gender valueobject.Gender) {
	c.gender = gender
}

func (c *Customer) SetBirthDate(birthDate valueobject.BirthDate) {
	c.birthDate = birthDate
}
