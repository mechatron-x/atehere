package port

import "github.com/mechatron-x/atehere/internal/core"

type (
	AllPaymentsDoneEventPublisher interface {
		NotifyEvent(event core.AllPaymentsDoneEvent)
	}

	Authenticator interface {
		GetUserID(idToken string) (string, error)
		GetUserEmail(idToken string) (string, error)
	}
)
