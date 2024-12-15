package port

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	EventNotifier interface {
		NotifyOrderCreatedEvent(event *dto.NewOrderEvent) error
		NotifyCheckoutEvent(event *dto.CheckoutEvent) error
	}

	NewOrderEventPublisher interface {
		NotifyEvent(event core.NewOrderEvent)
	}

	CheckoutEventPublisher interface {
		NotifyEvent(event core.CheckoutEvent)
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
	}
)
