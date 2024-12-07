package consumer

import (
	"fmt"

	"github.com/mechatron-x/atehere/internal/core"
)

type CreateBillConsumer struct {
}

func CreateBill() *CreateBillConsumer {
	return &CreateBillConsumer{}
}

func (rcv *CreateBillConsumer) ProcessEvent(event core.SessionClosedEvent) error {
	fmt.Println(event)
	return nil
}
