package valueobject

import "errors"

type EventType int

const (
	DefaultEvent EventType = iota
	NewOrderPlaced
	CallWaiter
	CheckoutRequested
)

func ParseEventType(eventType int) (EventType, error) {
	switch eventType {
	case int(DefaultEvent):
		return -1, errors.New("default event is not an event type")
	case int(NewOrderPlaced):
		return NewOrderPlaced, nil
	case int(CallWaiter):
		return CallWaiter, nil
	case int(CheckoutRequested):
		return CheckoutRequested, nil
	default:
		return -1, errors.New("unsupported event type")
	}
}

func (e EventType) String() string {
	switch e {
	case DefaultEvent:
		return ""
	case NewOrderPlaced:
		return "has a new order"
	case CallWaiter:
		return "calls the waiter"
	case CheckoutRequested:
		return "requests the bill"
	default:
		return ""
	}
}

func (e EventType) Equals(eventType EventType) bool {
	return int(e) == int(eventType)
}
