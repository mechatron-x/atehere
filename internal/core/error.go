package core

import (
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
)

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
