package consumer

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/session/port"
)

type CloseSessionConsumer struct {
	sessionRepository port.SessionRepository
}

func NewCloseSession(
	sessionRepository port.SessionRepository,
) *CloseSessionConsumer {
	return &CloseSessionConsumer{
		sessionRepository: sessionRepository,
	}
}

func (rcv *CloseSessionConsumer) ProcessEvent(event core.AllPaymentsDoneEvent) error {
	session, err := rcv.sessionRepository.GetByID(event.SessionID())
	if err != nil {
		return err
	}

	err = session.Close()
	if err != nil {
		return err
	}

	return rcv.sessionRepository.Save(session)
}
