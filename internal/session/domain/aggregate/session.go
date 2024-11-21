package aggregate

import (
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
	orders    []entity.Order
	events    []entity.Event
}

func NewSession() *Session {
	return &Session{
		Aggregate: core.NewAggregate(),
		startTime: time.Now(),
		orders:    make([]entity.Order, 0),
		events:    make([]entity.Event, 0),
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

func (s *Session) Events() []entity.Event {
	return s.events
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

func (s *Session) SetEvents(events []entity.Event) {
	s.events = events
}

func (s *Session) AddOrders(orders ...entity.Order) {
	for _, o := range orders {
		s.addOrder(o)
	}
}

func (s *Session) addOrder(order entity.Order) {
	for _, o := range s.orders {
		if o.ID() == order.ID() {
			return
		}
	}

	s.addEvent(valueobject.NewOrderPlaced)
	s.orders = append(s.orders, order)
}

func (s *Session) addEvent(eventType valueobject.EventType) {
	for _, e := range s.events {
		if eventType.Equals(e.EventType()) {
			e.SetInvokeTime(time.Now())
			return
		}
	}

	newEvent := entity.NewEvent()
	newEvent.SetEventType(eventType)

	s.events = append(s.events, newEvent)
}
