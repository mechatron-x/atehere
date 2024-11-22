package aggregate

import (
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

func (s *Session) PlaceOrders(orders ...entity.Order) {
	for _, o := range orders {
		err := s.addOrder(o)
		if err != nil {
			continue
		}

		s.RaiseEvent(event.NewOrderCreated(s.tableID, o.ID()))
	}
}

func (s *Session) Close() {
	s.SetDeletedAt(time.Now())
	s.RaiseEvent(event.NewSessionClosed(s.tableID, s.orders...))
}

func (s *Session) addOrder(order entity.Order) error {
	for _, o := range s.orders {
		if o.ID() == order.ID() {
			return fmt.Errorf("order with id %s already placed", o.ID())
		}
	}
	return nil
}
