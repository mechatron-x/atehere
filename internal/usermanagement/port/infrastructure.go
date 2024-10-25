package port

import "errors"

var (
	ErrUserUnauthorized = errors.New("user unauthorized")
	ErrUserCreation     = errors.New("user creation failed")
)

type (
	Authenticator interface {
		CreateUser(id, email, password string) error
		RevokeRefreshTokens(idToken string) error
		GetUserID(idToken string) (string, error)
		GetUserEmail(idToken string) (string, error)
	}
)
