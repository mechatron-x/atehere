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
	events        chan core.DomainEvent
}

func NewSession(
	repository port.SessionRepository,
	authenticator port.Authenticator,
	eventPusher port.EventPusher,
	eventBusSize int,
) *Session {
	session := &Session{
		repository:    repository,
		authenticator: authenticator,
		eventPusher:   eventPusher,
		events:        make(chan core.DomainEvent, eventBusSize),
	}

	for i := 0; i < eventBusSize; i++ {
		session.processEventsAsync()
	}

	return session
}

func (ss *Session) PlaceOrder(idToken string, tableID string, placeOrder dto.PlaceOrder) error {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return core.NewUnauthorizedError(err)
	}
	placeOrder.OrderedBy = customerID

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	order, err := placeOrder.ToEntity()
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

	session.PlaceOrders(order)

	err = ss.repository.Save(session)
	if err != nil {
		return core.NewPersistenceFailureError(err)
	}

	ss.pushEventsAsync(session.Events())

	return nil
}

func (ss *Session) pushEventsAsync(events []core.DomainEvent) {
	go func(events []core.DomainEvent) {
		for _, event := range events {
			ss.events <- event
		}
	}(events)
}

func (ss *Session) processEventsAsync() {
	go func(eventChan <-chan core.DomainEvent) {

	}(ss.events)
}
