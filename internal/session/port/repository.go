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
		GetTableOrdersView(tableID uuid.UUID) ([]dto.TableOrderView, error)
		GetManagerOrdersView(tableID uuid.UUID) ([]dto.ManagerOrderView, error)
		CheckoutEventView(sessionID uuid.UUID) (*dto.CheckoutEvent, error)
	}
)
