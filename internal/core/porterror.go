package core

type PortError interface {
	Context() string
	Reason() error
	Error() string
}

type BasePortError struct {
	context string
	reason  error
}

func (p *BasePortError) Context() string {
	return p.context
}

func (p *BasePortError) Reason() error {
	return p.reason
}

func (p *BasePortError) Error() string {
	return p.reason.Error()
}

type (
	//General Errors
	ConnectionError struct {
		BasePortError
	}

	// Repository Errors
	DataMappingError struct {
		BasePortError
	}

	DataNotFoundError struct {
		BasePortError
	}

	DataConflictError struct {
		BasePortError
	}

	DataPersistenceError struct {
		BasePortError
	}

	// Authenticator Errors
	AuthenticationFailedError struct {
		BasePortError
	}

	AuthorizationFailedError struct {
		BasePortError
	}

	TokenInvalidError struct {
		BasePortError
	}

	TokenExpiredError struct {
		BasePortError
	}

	DataValidationError struct {
		BasePortError
	}
)

func NewConnectionError(context string, reason error) *ConnectionError {
	return &ConnectionError{
		BasePortError{context: context, reason: reason},
	}
}

func NewDataMappingError(context string, reason error) *DataMappingError {
	return &DataMappingError{
		BasePortError{context: context, reason: reason},
	}
}

func NewDataNotFoundError(context string, reason error) *DataNotFoundError {
	return &DataNotFoundError{
		BasePortError{context: context, reason: reason},
	}
}

func NewDataConflictError(context string, reason error) *DataConflictError {
	return &DataConflictError{
		BasePortError{context: context, reason: reason},
	}
}

func NewDataPersistenceError(context string, reason error) *DataPersistenceError {
	return &DataPersistenceError{
		BasePortError{context: context, reason: reason},
	}
}

func NewAuthenticationFailedError(context string, reason error) *AuthenticationFailedError {
	return &AuthenticationFailedError{
		BasePortError{context: context, reason: reason},
	}
}

func NewAuthorizationFailedError(context string, reason error) *AuthorizationFailedError {
	return &AuthorizationFailedError{
		BasePortError{context: context, reason: reason},
	}
}

func NewTokenInvalidError(context string, reason error) *TokenInvalidError {
	return &TokenInvalidError{
		BasePortError{context: context, reason: reason},
	}
}

func NewTokenExpiredError(context string, reason error) *TokenExpiredError {
	return &TokenExpiredError{
		BasePortError{context: context, reason: reason},
	}
}

func NewDataValidationError(context string, reason error) *DataValidationError {
	return &DataValidationError{
		BasePortError{context: context, reason: reason},
	}
}
