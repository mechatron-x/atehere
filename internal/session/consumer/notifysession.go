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

func (rcv *NotifySessionConsumer) ProcessEvent(event core.SessionClosedEvent) error {
	sessionClosedEvent, err := rcv.sessionViewRepository.SessionClosedEventView(event.SessionID())
	if err != nil {
		return err
	}

	sessionClosedEvent.InvokeTime = event.InvokeTime().Unix()
	sessionClosedEvent.ID = event.ID()
	err = rcv.eventNotifier.NotifySessionClosedEvent(sessionClosedEvent)
	if err != nil {
		return err
	}

	return nil
}
