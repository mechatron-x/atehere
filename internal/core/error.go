package core

var (
	CodeUnauthorized       = "domain.error.unauthorized"
	CodeResourceNotFound   = "domain.error.resource_not_found"
	CodeDataConflict       = "domain.error.data_conflict"
	CodePersistenceFailure = "domain.error.persistence_failure"
	CodeValidationFailure  = "domain.error.validation_failure"
)

// DomainError interface defines the structure for all domain errors.
type DomainError interface {
	Code() string
	Error() string
}

// baseDomainError provides a common structure for domain errors.
type baseDomainError struct {
	code   string
	reason error
}

func (d baseDomainError) Code() string {
	return d.code
}

func (d baseDomainError) Error() string {
	return d.reason.Error()
}

type UnauthorizedError struct {
	baseDomainError
}

func NewUnauthorizedError(reason error) *UnauthorizedError {
	return &UnauthorizedError{
		baseDomainError: baseDomainError{
			code:   CodeUnauthorized,
			reason: reason,
		},
	}
}

type ResourceNotFoundError struct {
	baseDomainError
}

func NewResourceNotFoundError(reason error) *ResourceNotFoundError {
	return &ResourceNotFoundError{
		baseDomainError: baseDomainError{
			code:   CodeResourceNotFound,
			reason: reason,
		},
	}
}

type DataConflictError struct {
	baseDomainError
}

func NewDataConflictError(reason error) *DataConflictError {
	return &DataConflictError{
		baseDomainError: baseDomainError{
			code:   CodeDataConflict,
			reason: reason,
		},
	}
}

type PersistenceFailureError struct {
	baseDomainError
}

func NewPersistenceFailureError(reason error) *PersistenceFailureError {
	return &PersistenceFailureError{
		baseDomainError: baseDomainError{
			code:   CodePersistenceFailure,
			reason: reason,
		},
	}
}

type ValidationFailureError struct {
	baseDomainError
}

func NewValidationFailureError(reason error) *ValidationFailureError {
	return &ValidationFailureError{
		baseDomainError: baseDomainError{
			code:   CodeValidationFailure,
			reason: reason,
		},
	}
}
