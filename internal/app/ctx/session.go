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
	newOrderEventPublisher *broker.Publisher[core.NewOrderEvent],
	checkoutEventPublisher *broker.Publisher[core.CheckoutEvent],
) Session {
	repo := repository.NewSession(db)
	viewRepo := repository.NewSessionView(db)

	service := service.NewSession(
		repo,
		viewRepo,
		authenticator,
		newOrderEventPublisher,
		checkoutEventPublisher,
	)

	handler := handler.NewSession(*service)
	return Session{
		handler: handler,
	}
}

func (s Session) Handler() handler.SessionHandler {
	return s.handler
}

func NewNotifyOrderConsumer(
	db *gorm.DB,
	eventNotifier port.EventNotifier,
) broker.Consumer[core.NewOrderEvent] {
	viewRepo := repository.NewSessionView(db)
	notifyOrderConsumer := consumer.NewNotifyOrder(viewRepo, eventNotifier)

	return notifyOrderConsumer
}

func NewCheckoutConsumer(
	db *gorm.DB,
	eventNotifier port.EventNotifier,
) broker.Consumer[core.CheckoutEvent] {
	viewRepo := repository.NewSessionView(db)
	notifyCheckout := consumer.NewNotifyCheckout(viewRepo, eventNotifier)

	return notifyCheckout
}

func NewSessionClosedConsumer(
	db *gorm.DB,
) broker.Consumer[core.AllPaymentsDoneEvent] {
	repository := repository.NewSession(db)

	return consumer.NewCloseSession(repository)
}
