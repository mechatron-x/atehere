package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/billing/dto"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"gorm.io/gorm"
)

type BillRepository struct {
	db     *gorm.DB
	mapper mapper.BillMapper
}

func NewBill(db *gorm.DB) *BillRepository {
	return &BillRepository{
		db:     db,
		mapper: mapper.NewBill(),
	}
}

func (rcv *BillRepository) GetBySessionID(sessionID uuid.UUID) (*aggregate.Bill, error) {
	billModel := new(model.Bill)

	result := rcv.db.
		Preload("BillItems.Payments").
		Where(&model.Bill{SessionID: sessionID.String()}).
		First(billModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return rcv.mapper.FromModel(billModel)
}

func (rcv *BillRepository) Save(bill *aggregate.Bill) error {
	billModel := rcv.mapper.FromAggregate(bill)

	tx := rcv.db.Begin()
	defer tx.Commit()

	result := tx.First(&model.Bill{ID: billModel.ID})
	if result.RowsAffected == 0 {
		result = tx.Create(billModel)

		if result.Error != nil {
			tx.Rollback()
		}

		return result.Error
	}

	err := tx.Model(billModel).
		Association("BillItems").
		Unscoped().
		Replace(billModel.BillItems)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(billModel)
	if result.Error != nil {
		tx.Rollback()
	}

	return result.Error
}

type BillViewRepository struct {
	db *gorm.DB
}

func NewBillView(db *gorm.DB) *BillViewRepository {
	return &BillViewRepository{
		db: db,
	}
}

func (rcv *BillViewRepository) GetPostOrders(sessionID uuid.UUID) ([]dto.PostOrder, error) {
	var orders []dto.PostOrder

	result := rcv.db.Table("post_orders").
		Where("session_id = ?", sessionID.String()).
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
