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
		OrderCreatedEventView(sessionID, orderID uuid.UUID) (*dto.NewOrderEvent, error)
		GetTableOrdersView(sessionID uuid.UUID) ([]dto.TableOrderView, error)
		GetManagerOrdersView(sessionID uuid.UUID) ([]dto.ManagerOrderView, error)
		CheckoutEventView(sessionID uuid.UUID) (*dto.CheckoutEvent, error)
	}
)
