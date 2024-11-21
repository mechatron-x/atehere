package port

import "github.com/mechatron-x/atehere/internal/session/domain/aggregate"

type (
	EventPusher interface {
		Push(session *aggregate.Session) error
	}
	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
