package service

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/domain/event"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/session/port"
	"go.uber.org/zap"
)

type SessionService struct {
	repository                  port.SessionRepository
	viewRepository              port.SessionViewRepository
	authenticator               port.Authenticator
	orderCreatedEventPublisher  port.OrderCreatedEventPublisher
	sessionClosedEventPublisher port.SessionClosedEventPublisher
	log                         *zap.Logger
}

func NewSession(
	repository port.SessionRepository,
	viewRepository port.SessionViewRepository,
	authenticator port.Authenticator,
	orderCreatedEventPublisher port.OrderCreatedEventPublisher,
	sessionClosedEventPublisher port.SessionClosedEventPublisher,
) *SessionService {
	session := &SessionService{
		repository:                  repository,
		viewRepository:              viewRepository,
		authenticator:               authenticator,
		orderCreatedEventPublisher:  orderCreatedEventPublisher,
		sessionClosedEventPublisher: sessionClosedEventPublisher,
		log:                         logger.Instance(),
	}

	return session
}

func (ss *SessionService) PlaceOrders(idToken, tableID string, placeOrders *dto.PlaceOrders) error {
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

func (ss *SessionService) CustomerOrdersView(idToken, tableID string) (*dto.OrderList, error) {
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
	if len(orders) != 0 {
		ordersView.Currency = orders[0].Currency
	}

	return ordersView, nil
}

func (ss *SessionService) ManagerOrdersView(idToken, tableID string) (*dto.OrderList, error) {
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
	if len(managerOrders) != 0 {
		ordersView.Currency = managerOrders[0].Currency
	}

	return ordersView, nil
}

func (ss *SessionService) TableOrdersView(tableID string) (*dto.OrderList, error) {
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

func (ss *SessionService) Checkout(idToken, tableID string) error {
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

func (ss *SessionService) getActiveSession(tableID uuid.UUID) *aggregate.Session {
	session, err := ss.repository.GetByTableID(tableID)
	if err != nil {
		session = aggregate.NewSession()
		session.SetTableID(tableID)
		return session
	}

	return session
}

func (ss *SessionService) pushEventsAsync(events []core.DomainEvent) {
	go func(events []core.DomainEvent) {
		for _, e := range events {
			if orderCreatedEvent, ok := e.(event.OrderCreated); ok {
				ss.orderCreatedEventPublisher.NotifyEvent(orderCreatedEvent)
			} else if sessionClosedEvent, ok := e.(core.SessionClosedEvent); ok {
				ss.sessionClosedEventPublisher.NotifyEvent(sessionClosedEvent)
			} else {
				ss.log.Warn("unsupported event type skipping event processing")
			}
		}
	}(events)
}
