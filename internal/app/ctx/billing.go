package ctx

import (
	"github.com/mechatron-x/atehere/internal/billing/consumer"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/infrastructure/broker"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/repository"
	"gorm.io/gorm"
)

type PostOrderCtx struct {
}

func NewPostOrder(db *gorm.DB, sessionClosedPublisher *broker.Publisher[core.CheckoutEvent]) PostOrderCtx {
	postOrdersRepository := repository.NewPostOrder(db)
	createPostOrdersConsumer := consumer.NewCreatePostOrders(postOrdersRepository)

	sessionClosedPublisher.AddConsumer(createPostOrdersConsumer)
	return PostOrderCtx{}
}
