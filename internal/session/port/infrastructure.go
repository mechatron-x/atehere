package port

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	EventNotifier interface {
		NotifyOrderCreatedEvent(event *dto.OrderCreatedEvent) error
		NotifySessionClosedEvent(event *dto.SessionClosedEvent) error
	}

	OrderCreatedEventPublisher interface {
		NotifyEvent(event core.OrderCreatedEvent)
	}

	SessionClosedEventPublisher interface {
		NotifyEvent(event core.SessionClosedEvent)
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
