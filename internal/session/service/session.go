package service

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/session/port"
)

type Session struct {
	repository    port.SessionRepository
	authenticator port.Authenticator
	eventPusher   port.EventPusher
}

func NewSession(
	repository port.SessionRepository,
	authenticator port.Authenticator,
	eventPusher port.EventPusher,
) *Session {
	return &Session{
		repository:    repository,
		authenticator: authenticator,
		eventPusher:   eventPusher,
	}
}

func (ss *Session) PlaceOrder(idToken string, placeOrders dto.PlaceOrders) error {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return core.NewUnauthorizedError(err)
	}

	verifiedTableID, err := uuid.Parse(placeOrders.TableID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	orders, err := placeOrders.ToEntities(customerID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	hasSession := ss.repository.HasActiveSessions(verifiedTableID)
	var session *aggregate.Session

	if hasSession {
		session, err = ss.repository.GetByTableID(verifiedTableID)
		if err != nil {
			return core.NewResourceNotFoundError(err)
		}
	} else {
		session := aggregate.NewSession()
		session.SetTableID(verifiedTableID)
	}

	session.AddOrders(orders...)

	err = ss.repository.Save(session)
	if err != nil {
		return core.NewPersistenceFailureError(err)
	}

	go ss.eventPusher.Push(session)

	return nil
}

//TODO:
