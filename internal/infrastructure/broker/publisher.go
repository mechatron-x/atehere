package broker

import (
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type Publisher[TEvent Event] struct {
	consumers []Consumer[TEvent]
	log       *zap.Logger
}

func NewPublisher[TEvent Event]() *Publisher[TEvent] {
	return &Publisher[TEvent]{
		consumers: make([]Consumer[TEvent], 0),
		log:       logger.Instance(),
	}
}

func (rcv *Publisher[TEvent]) AddConsumer(consumer ...Consumer[TEvent]) {
	rcv.consumers = append(rcv.consumers, consumer...)
}

func (rcv *Publisher[TEvent]) NotifyEvent(event TEvent) {
	for _, consumer := range rcv.consumers {
		go func(consumer Consumer[TEvent], event TEvent) {
			err := consumer.ProcessEvent(event)
			if err != nil {
				rcv.log.Warn("Cannot process event", zap.Any("reason", err))
			}
		}(consumer, event)
	}
}
