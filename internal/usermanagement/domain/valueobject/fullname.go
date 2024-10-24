package valueobject

import (
	"errors"
	"fmt"
	"strings"
)

type FullName struct {
	firstName  string
	middleName string
	lastName   string
}

func NewFullName(name string) (FullName, error) {
	splitName := strings.Fields(name)
	if len(splitName) == 0 {
		return FullName{}, errors.New("name should not be empty")
	}

	fullName := FullName{}
	fullName.firstName = splitName[0]

	if len(splitName) == 2 {
		fullName.lastName = splitName[1]
	}

	if len(splitName) > 2 {
		fullName.middleName = splitName[1]
		fullName.lastName = splitName[2]
	}

	return fullName, nil
}

func (fn FullName) FirstName() string {
	return fn.firstName
}

func (fn FullName) MiddleName() string {
	return fn.middleName
}

func (fn FullName) LastName() string {
	return fn.lastName
}

func (fn FullName) String() string {
	return fmt.Sprintf("%s %s %s", fn.firstName, fn.middleName, fn.lastName)
}
