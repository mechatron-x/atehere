package port

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	EventNotifier interface {
		NotifyOrderCreatedEvent(event *dto.OrderCreatedEvent) error
		NotifyCheckoutEvent(event *dto.CheckoutEvent) error
	}

	OrderCreatedEventPublisher interface {
		NotifyEvent(event core.OrderCreatedEvent)
	}

	SessionClosedEventPublisher interface {
		NotifyEvent(event core.CheckoutEvent)
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
