package aggregate

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/event"
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

	for _, o := range orders {
		if err := s.placeOrderPolicy(o); err != nil {
			errs = append(errs, err)
			continue
		}

		s.orders = append(s.orders, o)
		s.RaiseEvent(event.NewOrderCreated(s.ID(), o.ID(), o.Quantity().Int()))
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (s *Session) Close(customerID uuid.UUID) error {
	err := s.closeSessionPolicy(customerID)
	if err != nil {
		return err
	}

	s.SetDeletedAt(time.Now())
	s.endTime = time.Now()
	s.RaiseEvent(event.NewSessionClosed(s.ID(), s.orders...))
	s.orders = make([]entity.Order, 0)

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

func (s *Session) closeSessionPolicy(customerID uuid.UUID) error {
	for _, order := range s.orders {
		if order.OrderedBy() == customerID {
			return nil
		}
	}

	return errors.New("the session cannot be closed because, specified customer is not a participant")
}
