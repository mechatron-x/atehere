package mapper

import (
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"gorm.io/gorm"
)

type PostOrderMapper struct{}

func (rcv PostOrderMapper) FromAggregate(postOrder *aggregate.PostOrder) *model.PostOrder {
	return &model.PostOrder{
		ID:         postOrder.ID().String(),
		SessionID:  postOrder.SessionID().String(),
		OrderedBy:  postOrder.OrderedBy().String(),
		MenuItemID: postOrder.MenuItemID().String(),
		Quantity:   postOrder.Quantity().Int(),
		Model: gorm.Model{
			CreatedAt: postOrder.CreatedAt(),
			UpdatedAt: postOrder.UpdatedAt(),
		},
	}
}
