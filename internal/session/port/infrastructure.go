package port

import (
	"time"

	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	EventPusher interface {
		PushOrderCreatedEvent(event *dto.OrderCreatedEventView, invokeTime time.Time) error
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
