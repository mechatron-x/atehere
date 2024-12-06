package port

import (
	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	EventNotifier interface {
		NotifyOrderCreatedEvent(event *dto.OrderCreatedEvent) error
		NotifySessionClosedEvent(event *dto.SessionClosedEvent) error
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
