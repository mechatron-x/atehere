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
	}

	SessionViewRepository interface {
		OrderCreatedEventView(sessionID, orderID uuid.UUID) (*dto.OrderCreatedEvent, error)
		SessionClosedEventView(sessionID uuid.UUID) (*dto.SessionClosedEvent, error)
		CustomerOrdersView(customerID, tableID uuid.UUID) ([]dto.Order, error)
		ManagerOrdersView(tableID uuid.UUID) ([]dto.Order, error)
		GetCustomersInSession(tableID uuid.UUID) ([]dto.SessionCustomer, error)
	}
)
