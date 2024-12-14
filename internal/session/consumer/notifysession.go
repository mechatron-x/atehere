package consumer

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/port"
)

type NotifySessionConsumer struct {
	sessionViewRepository port.SessionViewRepository
	eventNotifier         port.EventNotifier
}

func NewNotifySession(
	sessionViewRepository port.SessionViewRepository,
	eventNotifier port.EventNotifier,
) *NotifySessionConsumer {
	return &NotifySessionConsumer{
		sessionViewRepository: sessionViewRepository,
		eventNotifier:         eventNotifier,
	}
}

func (rcv *NotifySessionConsumer) ProcessEvent(event core.CheckoutEvent) error {
	checkoutEvent, err := rcv.sessionViewRepository.CheckoutEventView(event.SessionID())
	if err != nil {
		return err
	}

	checkoutEvent.InvokeTime = event.InvokeTime().Unix()
	checkoutEvent.ID = event.ID()
	err = rcv.eventNotifier.NotifyCheckoutEvent(checkoutEvent)
	if err != nil {
		return err
	}

	return nil
}
