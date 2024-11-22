package infrastructure

import (
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"go.uber.org/zap"
)

type FirebaseEventNotifier struct {
	log *zap.Logger
}

func NewFirebaseEventNotifier() (*FirebaseEventNotifier, error) {
	return &FirebaseEventNotifier{
		log: logger.Instance(),
	}, nil
}

func (fen *FirebaseEventNotifier) NotifyOrderCreatedEvent(event *dto.OrderCreatedEventView) error {
	fen.log.Info(event.Message())
	return nil
}
