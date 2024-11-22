package event

import (
	"slices"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
)

type SessionClosed struct {
	core.BaseEvent
	sessionID uuid.UUID
	orders    []entity.Order
}

func NewSessionClosed(sessionID uuid.UUID, orders ...entity.Order) SessionClosed {
	return SessionClosed{
		BaseEvent: core.NewDomainEvent(),
		sessionID: sessionID,
		orders:    slices.Clone(orders),
	}
}
