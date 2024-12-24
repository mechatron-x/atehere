package ctx

import (
	"github.com/mechatron-x/atehere/internal/billing/consumer"
	"github.com/mechatron-x/atehere/internal/billing/port"
	"github.com/mechatron-x/atehere/internal/billing/service"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/handler"
	"github.com/mechatron-x/atehere/internal/infrastructure/broker"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"gorm.io/gorm"
)

type BillingCtx struct {
	handler handler.BillingHandler
}

func NewBilling(db *gorm.DB, authenticator port.Authenticator, allPaymentsDonePublisher port.AllPaymentsDoneEventPublisher) BillingCtx {
	billingRepo := repository.NewBill(db)
	billingViewRepo := repository.NewBillView(db)
	service := service.NewBilling(
		authenticator,
		billingRepo,
		billingViewRepo,
		allPaymentsDonePublisher,
	)

	handler := handler.NewBilling(service)

	return BillingCtx{
		handler: handler,
	}
}

func (rcv BillingCtx) Handler() handler.BillingHandler {
	return rcv.handler
}

func NewCreateBillConsumer(db *gorm.DB) broker.Consumer[core.CheckoutEvent] {
	billingRepository := repository.NewBill(db)
	billingViewRepository := repository.NewBillView(db)
	return consumer.NewCreateBill(billingRepository, billingViewRepository)
}
