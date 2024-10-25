package core

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrPersistence  = errors.New("persistence error")
	ErrValidation   = errors.New("validation error")
)

func MapPortErrorToDomain(portErr PortError) error {
	switch portErr.(type) {
	case *ConnectionError:
		return ErrPersistence
	case *DataMappingError:
		return ErrValidation
	case *DataNotFoundError:
		return ErrNotFound
	case *DataConflictError:
		return ErrConflict
	case *DataPersistenceError:
		return ErrPersistence
	case *AuthenticationFailedError:
		return ErrUnauthorized
	case *AuthorizationFailedError:
		return ErrUnauthorized
	case *TokenInvalidError:
		return ErrUnauthorized
	case *TokenExpiredError:
		return ErrUnauthorized
	default:
		return ErrPersistence
	}
}
