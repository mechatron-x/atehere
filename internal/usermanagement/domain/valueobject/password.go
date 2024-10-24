package valueobject

import (
	"errors"

	"github.com/mechatron-x/atehere/internal/core"
)

type Password struct {
	password string
}

func NewPassword(password string) (Password, error) {
	if core.IsEmptyString(password) {
		return Password{}, errors.New("empty password")
	}

	return Password{
		password: password,
	}, nil
}

func (p Password) String() string {
	return p.password
}
