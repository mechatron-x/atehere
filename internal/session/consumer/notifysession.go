package consumer

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/port"
)

type NotifyCheckoutConsumer struct {
	sessionViewRepository port.SessionViewRepository
	eventNotifier         port.EventNotifier
}

func NewNotifyCheckout(
	sessionViewRepository port.SessionViewRepository,
	eventNotifier port.EventNotifier,
) *NotifyCheckoutConsumer {
	return &NotifyCheckoutConsumer{
		sessionViewRepository: sessionViewRepository,
		eventNotifier:         eventNotifier,
	}
}

func (rcv *NotifyCheckoutConsumer) ProcessEvent(event core.CheckoutEvent) error {
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
