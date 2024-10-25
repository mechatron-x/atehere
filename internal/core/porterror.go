package core

import "errors"

type (
	PortError interface {
		Context() string
		Reason() error
		Error() string
	}
)

var (
	ErrUserUnauthorized = errors.New("user unauthorized")
	ErrUserCreation     = errors.New("user creation failed")
)

var (
	ErrModelMappingFailed = errors.New("mapping failed")
	ErrModelNotFound      = errors.New("not found")
	ErrModelAlreadyExists = errors.New("already exists")
	ErrDbConnection       = errors.New("database connection failed")
)
