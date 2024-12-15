package mapper

import (
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"gorm.io/gorm"
)

type Bill struct {
	bi BillItem
}

func NewBill() Bill {
	return Bill{
		bi: NewBillItem(),
	}
}

func (b Bill) FromModel(model *model.Bill) (*aggregate.Bill, error) {
	builder := aggregate.NewBillBuilder()
	builder.SetID(model.ID)
	builder.SetSessionID(model.SessionID)

	billItems, err := b.bi.FromModels(model.BillItems)
	if err != nil {
		return nil, err
	}
	builder.SetBillItems(billItems)
	builder.SetCreatedAt(model.CreatedAt)
	builder.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		builder.SetDeletedAt(model.DeletedAt.Time)
	}

	return builder.Build()
}

func (b Bill) FromAggregate(aggregate *aggregate.Bill) *model.Bill {
	billItems := b.bi.FromEntities(aggregate.ID(), aggregate.BillItems())

	return &model.Bill{
		ID:        aggregate.ID().String(),
		SessionID: aggregate.SessionID().String(),
		BillItems: billItems,
		Model: gorm.Model{
			CreatedAt: aggregate.CreatedAt(),
			UpdatedAt: aggregate.UpdatedAt(),
			DeletedAt: gorm.DeletedAt{
				Time:  aggregate.DeletedAt(),
				Valid: aggregate.IsDeleted(),
			},
		},
	}
}
