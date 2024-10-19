package core

import (
	"fmt"
)

type (
	DomainError interface {
		Model() string
		Message() string
		Error() string
	}
)

type (
	ModelValidationError struct {
		model   string
		message string
	}

	ModelCreationError struct {
		model   string
		message string
	}

	ModelPersistenceError struct {
		model   string
		message string
	}
)

func ErrModelValidation(model string, message error) DomainError {
	return &ModelValidationError{
		model:   model,
		message: message.Error(),
	}
}

func (e *ModelValidationError) Model() string {
	return e.model
}

func (e *ModelValidationError) Message() string {
	return e.message
}

func (e *ModelValidationError) Error() string {
	return fmt.Sprintf("%s model is not processable, reason: %s", e.model, e.message)
}

func ErrModelCreation(model string, message error) DomainError {
	return &ModelCreationError{
		model:   model,
		message: message.Error(),
	}
}

func (e *ModelCreationError) Model() string {
	return e.model
}

func (e *ModelCreationError) Message() string {
	return e.message
}

func (e *ModelCreationError) Error() string {
	return fmt.Sprintf("%s model creation failed, reason: %s", e.model, e.message)
}

func ErrModelPersistence(model string, message error) DomainError {
	return &ModelPersistenceError{
		model:   model,
		message: message.Error(),
	}
}

func (e *ModelPersistenceError) Model() string {
	return e.model
}

func (e *ModelPersistenceError) Message() string {
	return e.message
}

func (e *ModelPersistenceError) Error() string {
	return fmt.Sprintf("%s model persistence failed, reason: %s", e.model, e.message)
}
