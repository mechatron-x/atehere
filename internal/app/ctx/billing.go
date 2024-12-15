package ctx

import (
	"github.com/mechatron-x/atehere/internal/billing/consumer"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/infrastructure/broker"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"gorm.io/gorm"
)

func NewCreateBillConsumer(db *gorm.DB) broker.Consumer[core.CheckoutEvent] {
	billingRepository := repository.NewBill(db)
	billingViewRepository := repository.NewBillView(db)
	return consumer.NewCreateBill(billingRepository, billingViewRepository)
}
