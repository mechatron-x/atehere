package consumer

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/billing/port"
	"github.com/mechatron-x/atehere/internal/core"
)

type CreatePostOrdersConsumer struct {
	postOrderRepository port.PostOrderRepository
}

func NewCreatePostOrders(postOrderRepository port.PostOrderRepository) *CreatePostOrdersConsumer {
	return &CreatePostOrdersConsumer{
		postOrderRepository: postOrderRepository,
	}
}

func (rcv *CreatePostOrdersConsumer) ProcessEvent(event core.CheckoutEvent) error {
	postOrders, err := rcv.toPostOrders(event)
	if err != nil {
		return err
	}

	for _, po := range postOrders {
		err = rcv.postOrderRepository.Save(po)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rcv *CreatePostOrdersConsumer) toPostOrders(event core.CheckoutEvent) ([]*aggregate.PostOrder, error) {
	postOrders := make([]*aggregate.PostOrder, 0)
	for _, p := range event.Orders() {
		postOrder, err := rcv.toPostOrder(event.SessionID(), p)
		if err != nil {
			return nil, err
		}

		postOrders = append(postOrders, postOrder)
	}

	return postOrders, nil
}

func (rcv *CreatePostOrdersConsumer) toPostOrder(sessionID uuid.UUID, order core.Order) (*aggregate.PostOrder, error) {
	postOrder := aggregate.NewPostOrder()
	postOrder.SetSessionID(sessionID)
	postOrder.SetMenuItemID(order.MenuItemID())
	postOrder.SetOrderedBy(order.OrderedBy())

	verifiedQuantity, err := valueobject.NewQuantity(order.Quantity())
	if err != nil {
		return nil, err
	}
	postOrder.SetQuantity(verifiedQuantity)

	return postOrder, nil
}
