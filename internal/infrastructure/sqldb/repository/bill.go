package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/billing/dto"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/view"
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

	result := rcv.db.Table(view.PostOrders).
		Where("session_id = ?", sessionID.String()).
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (rcv *BillViewRepository) GetPastBills(customerID uuid.UUID) ([]dto.PastBill, error) {
	var pastBillItems []model.PastBillsView

	result := rcv.db.Table(view.PastBillItems).
		Where(&model.PastBillsView{OwnerID: customerID.String()}).
		Scan(&pastBillItems)
	if result.Error != nil {
		return nil, result.Error
	}

	pastBillsMap := make(map[string][]dto.PastBillItem)
	pastBills := make([]dto.PastBill, 0)
	for _, pb := range pastBillItems {
		pastBillItem := dto.PastBillItem{
			ItemName:   pb.ItemName,
			Quantity:   pb.Quantity,
			UnitPrice:  pb.UnitPrice,
			OrderPrice: pb.OrderPrice,
			PaidPrice:  pb.PaidPrice,
			Currency:   pb.Currency,
		}

		_, ok := pastBillsMap[pb.BillID]
		if !ok {
			pastBillsMap[pb.BillID] = make([]dto.PastBillItem, 0)
			pastBill := dto.PastBill{
				BillID:         pb.BillID,
				RestaurantName: pb.RestaurantName,
			}
			pastBills = append(pastBills, pastBill)
		}

		pastBillsMap[pb.BillID] = append(pastBillsMap[pb.BillID], pastBillItem)
	}

	for i, pb := range pastBills {
		pastBillItems, ok := pastBillsMap[pb.BillID]
		if !ok {
			continue
		}

		pastBills[i].BillItems = pastBillItems
	}

	return pastBills, nil
}
