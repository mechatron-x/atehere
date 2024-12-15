package consumer

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/port"
)

type NotifyOrderConsumer struct {
	sessionViewRepository port.SessionViewRepository
	eventNotifier         port.EventNotifier
}

func NotifyOrder(
	sessionViewRepository port.SessionViewRepository,
	eventNotifier port.EventNotifier,
) *NotifyOrderConsumer {
	return &NotifyOrderConsumer{
		sessionViewRepository: sessionViewRepository,
		eventNotifier:         eventNotifier,
	}
}

func (rcv *NotifyOrderConsumer) ProcessEvent(event core.NewOrderEvent) error {
	orderCreatedEventView, err := rcv.sessionViewRepository.OrderCreatedEventView(event.SessionID(), event.OrderID())
	if err != nil {
		return err
	}

	orderCreatedEventView.Quantity = event.Quantity()
	orderCreatedEventView.InvokeTime = event.InvokeTime().Unix()
	orderCreatedEventView.ID = event.ID()
	err = rcv.eventNotifier.NotifyOrderCreatedEvent(orderCreatedEventView)
	if err != nil {
		return err
	}

	return nil
}
