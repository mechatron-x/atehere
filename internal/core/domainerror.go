package core

import (
	"errors"
	"fmt"
)

type (
	DomainError interface {
		Error() string
		Reason() error
	}
)

type (
	ValidationError struct {
		context string
		reason  error
	}

	PersistenceError struct {
		context string
		reason  error
	}

	ConflictError struct {
		context string
		reason  error
	}

	AuthorizationError struct {
		context string
		reason  error
	}

	NotFoundError struct {
		context string
		reason  error
	}

	UnhandledError struct {
		context string
		reason  error
	}
)

func MapPortErrorToDomain(context string, err error) DomainError {
	if errors.Is(err, ErrDbConnection) {
		return NewPersistenceError(context, err)
	}
	if errors.Is(err, ErrModelAlreadyExists) {
		return NewConflictError(context, err)
	}
	if errors.Is(err, ErrModelNotFound) {
		return NewNotFoundError(context, err)
	}
	if errors.Is(err, ErrModelMappingFailed) {
		return NewValidationError(context, err)
	}
	if errors.Is(err, ErrUserCreation) {
		return NewPersistenceError(context, err)
	}
	if errors.Is(err, ErrUserUnauthorized) {
		return NewAuthorizationError(context, err)
	}
	return NewUnhandledError(context, err)
}

func NewValidationError(context string, reason error) *ValidationError {
	return &ValidationError{
		context: context,
		reason:  reason,
	}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.context, v.reason)
}

func (v *ValidationError) Reason() error {
	return v.reason
}

func NewPersistenceError(context string, reason error) *PersistenceError {
	return &PersistenceError{
		context: context,
		reason:  reason,
	}
}

func (v *PersistenceError) Error() string {
	return fmt.Sprintf("%s: %s", v.context, v.reason)
}

func (v *PersistenceError) Reason() error {
	return v.reason
}

func NewConflictError(context string, reason error) *ConflictError {
	return &ConflictError{
		context: context,
		reason:  reason,
	}
}

func (c *ConflictError) Error() string {
	return fmt.Sprintf("%s: %s", c.context, c.reason)
}

func (c *ConflictError) Reason() error {
	return c.reason
}

func NewAuthorizationError(context string, reason error) *AuthorizationError {
	return &AuthorizationError{
		context: context,
		reason:  reason,
	}
}

func (a *AuthorizationError) Error() string {
	return fmt.Sprintf("%s: %s", a.context, a.reason)
}

func (a *AuthorizationError) Reason() error {
	return a.reason
}

func NewNotFoundError(context string, reason error) *NotFoundError {
	return &NotFoundError{
		context: context,
		reason:  reason,
	}
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", n.context, n.reason)
}

func (n *NotFoundError) Reason() error {
	return n.reason
}

func NewUnhandledError(context string, reason error) *UnhandledError {
	return &UnhandledError{
		context: context,
		reason:  reason,
	}
}

func (u *UnhandledError) Error() string {
	return fmt.Sprintf("%s: %s", u.context, u.reason)
}

func (u *UnhandledError) Reason() error {
	return u.reason
}
