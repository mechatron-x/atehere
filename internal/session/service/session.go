package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/session/port"
	"go.uber.org/zap"
)

type SessionService struct {
	repository             port.SessionRepository
	viewRepository         port.SessionViewRepository
	authenticator          port.Authenticator
	newOrderEventPublisher port.NewOrderEventPublisher
	checkoutEventPublisher port.CheckoutEventPublisher
	log                    *zap.Logger
}

func NewSession(
	repository port.SessionRepository,
	viewRepository port.SessionViewRepository,
	authenticator port.Authenticator,
	newOrderEventPublisher port.NewOrderEventPublisher,
	checkoutEventPublisher port.CheckoutEventPublisher,
) *SessionService {
	session := &SessionService{
		repository:             repository,
		viewRepository:         viewRepository,
		authenticator:          authenticator,
		newOrderEventPublisher: newOrderEventPublisher,
		checkoutEventPublisher: checkoutEventPublisher,
		log:                    logger.Instance(),
	}

	return session
}

func (ss *SessionService) PlaceOrders(idToken, tableID string, placeOrders *dto.PlaceOrders) (*dto.Session, error) {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	orders, err := placeOrders.ToEntities(customerID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	session := ss.getActiveSession(verifiedTableID)
	fmt.Println(session.EndTime().Format(time.ANSIC))
	err = session.PlaceOrders(orders...)
	if err != nil {
		return nil, core.NewDomainIntegrityViolationError(err)
	}

	err = ss.repository.Save(session)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	ss.pushEventsAsync(session.Events())
	return &dto.Session{SessionID: session.ID().String()}, nil
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

func (ss *SessionService) Checkout(idToken, tableID string) (*dto.Session, error) {
	customerID, err := ss.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	verifiedCustomerID, err := uuid.Parse(customerID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	verifiedTableID, err := uuid.Parse(tableID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	session := ss.getActiveSession(verifiedTableID)
	err = session.Checkout(verifiedCustomerID)
	if err != nil {
		return nil, core.NewDomainIntegrityViolationError(err)
	}

	err = ss.repository.Save(session)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	ss.pushEventsAsync(session.Events())

	return &dto.Session{SessionID: session.ID().String()}, nil
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
			if newOrderEvent, ok := e.(core.NewOrderEvent); ok {
				ss.newOrderEventPublisher.NotifyEvent(newOrderEvent)
			} else if checkoutEvent, ok := e.(core.CheckoutEvent); ok {
				ss.checkoutEventPublisher.NotifyEvent(checkoutEvent)
			} else {
				ss.log.Warn("unsupported event type skipping event processing")
			}
		}
	}(events)
}
