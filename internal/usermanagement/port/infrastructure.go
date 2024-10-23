package port

import (
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type (
	CustomerAuthRecord struct {
		UID           string
		Disabled      bool
		EmailVerified bool
		DisplayName   string
		Email         string
		PhoneNumber   string
	}

	CustomerAuthenticator interface {
		CreateUser(user *aggregate.Customer) error
		RevokeRefreshTokens(idToken string) error
		VerifyUser(idToken string) (*CustomerAuthRecord, error)
	}
)
