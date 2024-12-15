package aggregate

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
)

type Session struct {
	core.Aggregate
	tableID   uuid.UUID
	startTime time.Time
	endTime   time.Time
	state     valueobject.SessionState
	orders    []entity.Order
}

func NewSession() *Session {
	return &Session{
		Aggregate: core.NewAggregate(),
		startTime: time.Now(),
		state:     valueobject.Active,
		orders:    make([]entity.Order, 0),
	}
}

func (s *Session) TableID() uuid.UUID {
	return s.tableID
}

func (s *Session) StartTime() time.Time {
	return s.startTime
}

func (s *Session) EndTime() time.Time {
	return s.endTime
}

func (s *Session) State() valueobject.SessionState {
	return s.state
}

func (s *Session) Orders() []entity.Order {
	return s.orders
}

func (s *Session) SetTableID(tableID uuid.UUID) {
	s.tableID = tableID
}

func (s *Session) SetStartTime(startTime time.Time) {
	s.startTime = startTime
}

func (s *Session) SetEndTime(endTime time.Time) {
	s.endTime = endTime
}

func (s *Session) SetState(state valueobject.SessionState) {
	s.state = state
}

func (s *Session) SetOrders(orders []entity.Order) {
	s.orders = orders
}

func (s *Session) PlaceOrders(orders ...entity.Order) error {
	for _, o := range orders {
		if err := s.placeOrderPolicy(o); err != nil {
			return err
		}

		if err := s.placeOrder(o); err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) Checkout(customerID uuid.UUID) error {
	err := s.checkoutPolicy(customerID)
	if err != nil {
		return err
	}

	s.endTime = time.Now()
	s.state = valueobject.CheckoutPending

	s.RaiseEvent(core.NewCheckoutEvent(s.ID(), s.toEventOrders()))
	return nil
}

func (s *Session) Close() error {
	if err := s.closePolicy(); err != nil {
		return err
	}

	s.SetDeletedAt(time.Now())
	s.SetOrders([]entity.Order{})
	s.state = valueobject.Completed
	return nil
}

func (s *Session) placeOrder(newOrder entity.Order) error {
	if i, err := s.findPreviousOrder(newOrder.MenuItemID(), newOrder.OrderedBy()); err == nil {
		order := s.orders[i]
		if err := order.AddQuantity(newOrder.Quantity()); err != nil {
			return err
		}
		s.orders[i] = order
		s.RaiseEvent(core.NewOrderCreatedEvent(s.ID(), order.ID(), newOrder.Quantity().Int()))
		return nil
	}

	s.orders = append(s.orders, newOrder)
	s.RaiseEvent(core.NewOrderCreatedEvent(s.ID(), newOrder.ID(), newOrder.Quantity().Int()))
	return nil
}

func (s *Session) placeOrderPolicy(order entity.Order) error {
	if s.state != valueobject.Active {
		return fmt.Errorf("order cannot be placed: session is not in '%s' state", valueobject.Active)
	}

	for _, o := range s.orders {
		if o.ID() == order.ID() {
			return fmt.Errorf("order cannot be placed: order with id %s already exists", order.ID())
		}
	}

	return nil
}

func (s *Session) checkoutPolicy(customerID uuid.UUID) error {
	if s.state == valueobject.CheckoutPending {
		return errors.New("checkout cannot be proceed: checkout has already been requested for this session")
	}

	for _, order := range s.orders {
		if order.OrderedBy() == customerID {
			return nil
		}
	}

	return errors.New("checkout cannot be proceed: specified customer is not a participant")
}

func (s *Session) closePolicy() error {
	if s.state != valueobject.CheckoutPending {
		return errors.New("session cannot be closed: checkout needed")
	}

	return nil
}

func (s *Session) findPreviousOrder(menuItemID, orderedBy uuid.UUID) (int, error) {
	for i, o := range s.orders {
		if o.OrderedBy() != orderedBy || o.MenuItemID() != menuItemID {
			continue
		}

		return i, nil
	}

	return -1, fmt.Errorf("order with %s menu item id and %s customer id not found", menuItemID, orderedBy)
}

func (s *Session) toEventOrders() []core.Order {
	eventOrders := make([]core.Order, 0)
	for _, o := range s.orders {
		order := core.Order{
			MenuItemID: o.MenuItemID(),
			OrderedBy:  o.OrderedBy(),
			Quantity:   o.Quantity().Int(),
		}
		eventOrders = append(eventOrders, order)
	}

	return eventOrders
}
