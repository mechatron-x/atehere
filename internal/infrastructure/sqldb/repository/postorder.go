package repository

import (
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"gorm.io/gorm"
)

type PostOrderRepository struct {
	db     *gorm.DB
	mapper mapper.PostOrderMapper
}

func NewPostOrder(db *gorm.DB) *PostOrderRepository {
	return &PostOrderRepository{
		db:     db,
		mapper: mapper.PostOrderMapper{},
	}
}

func (rcv *PostOrderRepository) Save(postOrder *aggregate.PostOrder) error {
	postOrderModel := rcv.mapper.FromAggregate(postOrder)

	result := rcv.db.Save(postOrderModel)

	return result.Error
}
