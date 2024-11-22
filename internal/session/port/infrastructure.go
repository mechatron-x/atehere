package port

import (
	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	EventNotifier interface {
		NotifyOrderCreatedEvent(event *dto.OrderCreatedEventView) error
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
