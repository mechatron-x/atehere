package ctx

import (
	"github.com/mechatron-x/atehere/internal/billing/consumer"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/infrastructure/broker"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"gorm.io/gorm"
)

func NewCreatePostOrdersConsumer(db *gorm.DB) broker.Consumer[core.CheckoutEvent] {
	postOrdersRepository := repository.NewPostOrder(db)
	return consumer.NewCreatePostOrders(postOrdersRepository)
}
