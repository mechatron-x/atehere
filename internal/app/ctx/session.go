package ctx

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/httpserver/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/broker"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"github.com/mechatron-x/atehere/internal/session/consumer"
	"github.com/mechatron-x/atehere/internal/session/port"
	"github.com/mechatron-x/atehere/internal/session/service"
	"gorm.io/gorm"
)

type Session struct {
	handler handler.SessionHandler
}

func NewSession(
	db *gorm.DB,
	authenticator port.Authenticator,
	eventNotifier port.EventNotifier,
	orderCreatedEventPublisher *broker.Publisher[core.NewOrderEvent],
	sessionClosedEventPublisher *broker.Publisher[core.CheckoutEvent],
) Session {
	repo := repository.NewSession(db)
	viewRepo := repository.NewSessionView(db)

	notifyOrderCreatedEvent := consumer.NotifyOrder(viewRepo, eventNotifier)
	orderCreatedEventPublisher.AddConsumer(notifyOrderCreatedEvent)

	notifyConsumer := consumer.NewNotifyCheckout(viewRepo, eventNotifier)
	sessionClosedEventPublisher.AddConsumer(notifyConsumer)

	service := service.NewSession(
		repo,
		viewRepo,
		authenticator,
		orderCreatedEventPublisher,
		sessionClosedEventPublisher,
	)

	handler := handler.NewSession(*service)
	return Session{
		handler: handler,
	}
}

func (s Session) Handler() handler.SessionHandler {
	return s.handler
}
