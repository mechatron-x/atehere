package service

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/domain/event"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/session/port"
	"go.uber.org/zap"
)

type Session struct {
	repository     port.SessionRepository
	viewRepository port.SessionViewRepository
	authenticator  port.Authenticator
	eventNotifier  port.EventNotifier
	events         chan core.DomainEvent
	log            *zap.Logger
}

func NewSession(
	repository port.SessionRepository,
	viewRepository port.SessionViewRepository,
	authenticator port.Authenticator,
	eventPusher port.EventNotifier,
	eventBusSize int,
) *Session {
	session := &Session{
		repository:     repository,
		viewRepository: viewRepository,
		authenticator:  authenticator,
		eventNotifier:  eventPusher,
		events:         make(chan core.DomainEvent, eventBusSize),
		log:            logger.Instance(),
	}

	for i := 0; i < eventBusSize; i++ {
		session.processEventsAsync()
	}

	return session
}

func (ss *Session) PlaceOrders(idToken, tableID string, placeOrders *dto.PlaceOrders) error {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return core.NewUnauthorizedError(err)
	}

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	orders, err := placeOrders.ToEntities(customerID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	session := ss.getActiveSession(verifiedTableID)

	err = session.PlaceOrders(orders...)
	if err != nil {
		return core.NewDomainIntegrityViolationError(err)
	}

	err = ss.repository.Save(session)
	if err != nil {
		return core.NewPersistenceFailureError(err)
	}

	ss.pushEventsAsync(session.Events())

	return nil
}

func (ss *Session) CustomerOrdersView(idToken, tableID string) (*dto.OrderList, error) {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	orders, err := ss.viewRepository.GetTableOrdersView(verifiedTableID)
	if err != nil {
		return nil, err
	}

	ordersView := &dto.OrderList{
		Orders: dto.FromTableOrdersViewWithFilter(orders, func(tableOrder dto.TableOrderView) bool {
			return tableOrder.CustomerID == customerID
		}),
	}

	ordersView.CalculateTotalPrice()
	ordersView.CalculateTotalPrice()
	if len(orders) != 0 {
		ordersView.Currency = orders[0].Currency
	}

	return ordersView, nil
}

func (ss *Session) ManagerOrdersView(idToken, tableID string) (*dto.OrderList, error) {
	_, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	managerOrders, err := ss.viewRepository.GetManagerOrdersView(verifiedTableID)
	if err != nil {
		return nil, err
	}

	ordersView := &dto.OrderList{
		Orders: dto.FromManagerOrdersView(managerOrders),
	}

	ordersView.CalculateTotalPrice()
	ordersView.CalculateTotalPrice()
	if len(managerOrders) != 0 {
		ordersView.Currency = managerOrders[0].Currency
	}

	return ordersView, nil
}

func (ss *Session) TableOrdersView(tableID string) (*dto.OrderList, error) {
	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	tableOrders, err := ss.viewRepository.GetTableOrdersView(verifiedTableID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	ordersView := &dto.OrderList{
		Orders: dto.FromTableOrdersView(tableOrders),
	}

	ordersView.CalculateTotalPrice()
	if len(tableOrders) != 0 {
		ordersView.Currency = tableOrders[0].Currency
	}

	return ordersView, nil
}

func (ss *Session) Checkout(idToken, tableID string) error {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return core.NewUnauthorizedError(err)
	}

	verifiedCustomerID, err := uuid.Parse(customerID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	session := ss.getActiveSession(verifiedTableID)
	err = session.Close(verifiedCustomerID)
	if err != nil {
		return core.NewDomainIntegrityViolationError(err)
	}

	err = ss.repository.Save(session)
	if err != nil {
		return core.NewPersistenceFailureError(err)
	}

	ss.pushEventsAsync(session.Events())

	return nil
}

func (ss *Session) getActiveSession(tableID uuid.UUID) *aggregate.Session {
	session, err := ss.repository.GetByTableID(tableID)
	if err != nil {
		session = aggregate.NewSession()
		session.SetTableID(tableID)
		return session
	}

	return session
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
		for e := range eventChan {
			if orderCreatedEvent, ok := e.(event.OrderCreated); ok {
				_ = ss.processOrderCreatedEvent(orderCreatedEvent)
			} else if sessionClosedEvent, ok := e.(event.SessionClosed); ok {
				_ = ss.processSessionClosedEvent(sessionClosedEvent)
			} else {
				ss.log.Warn("unsupported event type skipping event processing")
			}
		}
	}(ss.events)
}

func (ss *Session) processOrderCreatedEvent(event event.OrderCreated) error {
	orderCreatedEventView, err := ss.viewRepository.OrderCreatedEventView(event.SessionID(), event.OrderID())
	if err != nil {
		return err
	}

	orderCreatedEventView.InvokeTime = event.InvokeTime().Unix()
	orderCreatedEventView.ID = event.ID()
	err = ss.eventNotifier.NotifyOrderCreatedEvent(orderCreatedEventView)
	if err != nil {
		return err
	}

	return nil
}

func (ss *Session) processSessionClosedEvent(event event.SessionClosed) error {
	sessionClosedEventView, err := ss.viewRepository.SessionClosedEventView(event.SessionID())
	if err != nil {
		return err
	}

	sessionClosedEventView.InvokeTime = event.InvokeTime().Unix()
	sessionClosedEventView.ID = event.ID()
	err = ss.eventNotifier.NotifySessionClosedEvent(sessionClosedEventView)
	if err != nil {
		return err
	}

	return nil
}
