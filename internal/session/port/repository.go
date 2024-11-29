package port

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/dto"
)

type (
	SessionRepository interface {
		Save(session *aggregate.Session) error
		GetByTableID(tableID uuid.UUID) (*aggregate.Session, error)
		HasActiveSessions(tableID uuid.UUID) bool
	}

	SessionViewRepository interface {
		OrderCreatedEventView(sessionID, orderID uuid.UUID) (*dto.OrderCreatedEventView, error)
		SessionClosedEventView(sessionID uuid.UUID) (*dto.SessionClosedEventView, error)
		OrderCustomerView(customerID uuid.UUID) ([]dto.OrderCustomerView, error)
		OrderTableView(tableID uuid.UUID) ([]dto.OrderTableView, error)
	}
)
