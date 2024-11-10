package valueobject

import (
	"errors"
	"regexp"

	"github.com/mechatron-x/atehere/internal/core"
)

type PhoneNumber struct {
	phoneNumber string
}

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)

func NewPhoneNumber(phoneNumber string) (PhoneNumber, error) {
	if core.IsEmptyString(phoneNumber) {
		return PhoneNumber{}, errors.New("phone number cannot be empty")
	}

	if !phoneRegex.MatchString(phoneNumber) {
		return PhoneNumber{}, errors.New("invalid phone number format")
	}

	return PhoneNumber{phoneNumber}, nil
}

func (p PhoneNumber) String() string {
	return p.phoneNumber
}
