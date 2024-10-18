package port

import (
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type (
	AuthUserRecord struct {
		UID           string
		Disabled      bool
		EmailVerified bool
		DisplayName   string
		Email         string
		PhoneNumber   string
	}

	AuthInfrastructure interface {
		CreateUser(user *aggregate.User) error
		RevokeRefreshTokens(idToken string) error
		VerifyUser(idToken string) (*AuthUserRecord, error)
	}
)
