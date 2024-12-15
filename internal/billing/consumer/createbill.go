package consumer

import (
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/port"
	"github.com/mechatron-x/atehere/internal/core"
)

type CreateBillConsumer struct {
	billRepository      port.BillRepository
	postOrderRepository port.BillViewRepository
}

func NewCreateBill(
	billRepository port.BillRepository,
	postOrderRepository port.BillViewRepository,
) *CreateBillConsumer {
	return &CreateBillConsumer{
		billRepository:      billRepository,
		postOrderRepository: postOrderRepository,
	}
}

func (rcv *CreateBillConsumer) ProcessEvent(event core.CheckoutEvent) error {
	postOrders, err := rcv.postOrderRepository.GetPostOrders(event.SessionID())
	if err != nil {
		return err
	}

	billItems := make([]entity.BillItem, 0)
	for _, po := range postOrders {
		billItem, err := po.ToBillItem()
		if err != nil {
			return err
		}

		billItems = append(billItems, *billItem)
	}

	billBuilder := aggregate.NewBillBuilder()
	billBuilder.SetSessionID(event.SessionID().String())
	billBuilder.SetBillItems(billItems)
	bill, err := billBuilder.Build()
	if err != nil {
		return err
	}

	return rcv.billRepository.Save(bill)
}
