package entity

import "github.com/mechatron-x/atehere/internal/core"

type Owner struct {
	core.Entity
	fullName string
	email    string
}

func NewOwner() *Owner {
	return &Owner{
		Entity: core.NewEntity(),
	}
}

func (o *Owner) FullName() string {
	return o.fullName
}

func (o *Owner) Email() string {
	return o.email
}

func (o *Owner) SetFullName(fullName string) {
	o.fullName = fullName
}

func (o *Owner) SetEmail(email string) {
	o.email = email
}
