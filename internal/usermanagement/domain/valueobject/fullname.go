package valueobject

import (
	"errors"
	"strings"
)

type FullName struct {
	firstName  string
	middleName string
	lastName   string
}

const (
	nameSeparator = " "
)

func NewFullName(name string) (FullName, error) {
	name = strings.TrimSpace(name)
	splitName := strings.Split(name, nameSeparator)
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
