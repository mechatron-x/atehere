package aggregate

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
)

type Session struct {
	core.Aggregate
	tableID   uuid.UUID
	startTime time.Time
	endTime   time.Time
	orders    []entity.Order
}

func NewSession() *Session {
	return &Session{
		Aggregate: core.NewAggregate(),
		startTime: time.Now(),
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

func (s *Session) SetOrders(orders []entity.Order) {
	s.orders = orders
}

func (s *Session) PlaceOrders(orders ...entity.Order) error {
	errs := make([]error, 0)

	if !s.endTime.IsZero() {
		return errors.New("checkout requested for this session")
	}

	for _, o := range orders {
		if err := s.placeOrder(o); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (s *Session) Checkout(customerID uuid.UUID) error {
	err := s.checkoutPolicy(customerID)
	if err != nil {
		return err
	}

	s.endTime = time.Now()
	s.RaiseEvent(core.NewCheckoutEvent(s.ID(), s.toEventOrders()))

	return nil
}

func (s *Session) placeOrder(newOrder entity.Order) error {
	if err := s.placeOrderPolicy(newOrder); err != nil {
		return err
	}

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
	for _, o := range s.orders {
		if o.ID() == order.ID() {
			return fmt.Errorf("order with id %s already exists", order.ID())
		}
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

func (s *Session) checkoutPolicy(customerID uuid.UUID) error {
	for _, order := range s.orders {
		if order.OrderedBy() == customerID {
			return nil
		}
	}

	return errors.New("the session cannot be closed because, specified customer is not a participant")
}

func (s *Session) toEventOrders() []core.Order {
	eventOrders := make([]core.Order, 0)
	for _, o := range s.orders {
		order := core.NewOrder(o.MenuItemID(), o.OrderedBy(), o.Quantity().Int())
		eventOrders = append(eventOrders, order)
	}

	return eventOrders
}
