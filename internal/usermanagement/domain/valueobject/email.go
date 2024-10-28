package valueobject

import (
	"errors"
	"regexp"
	"strings"

	"github.com/mechatron-x/atehere/internal/core"
)

const (
	emailSeparator = "@"
	emailRegexp    = "(?:[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+)*)@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,}"
)

type Email struct {
	username   string
	mailServer string
}

func NewEmail(email string) (Email, error) {
	if core.IsEmptyString(email) {
		return Email{}, errors.New("empty email address")
	}

	regexp, err := regexp.Compile(emailSeparator)
	if err != nil {
		return Email{}, err
	}

	if !regexp.MatchString(email) {
		return Email{}, errors.New("invalid mail format")
	}

	emailChunks := strings.Split(email, emailSeparator)
	return Email{
		username:   emailChunks[0],
		mailServer: emailChunks[1],
	}, nil
}

func (e Email) Username() string {
	return e.username
}

func (e Email) MailServer() string {
	return e.mailServer
}

func (e Email) String() string {
	return strings.Join([]string{e.username, e.mailServer}, emailSeparator)
}
