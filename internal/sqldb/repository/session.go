package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Session struct {
	db     *gorm.DB
	mapper mapper.Session
}

func NewSession(db *gorm.DB) *Session {
	return &Session{
		db:     db,
		mapper: mapper.Session{},
	}
}

func (s *Session) Save(session *aggregate.Session) error {
	sessionModel := s.mapper.FromAggregate(session)

	tx := s.db.Begin()
	defer tx.Commit()

	result := tx.First(&model.Session{ID: sessionModel.ID})

	if result.RowsAffected == 0 {
		result = tx.Create(sessionModel)

		if result.Error != nil {
			tx.Rollback()
		}

		return result.Error
	}

	err := tx.Model(sessionModel).
		Association("Orders").
		Unscoped().
		Replace(sessionModel.Orders)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(sessionModel)
	if result.Error != nil {
		tx.Rollback()
	}

	return result.Error
}

func (s *Session) GetByTableID(tableID uuid.UUID) (*aggregate.Session, error) {
	sessionModel := new(model.Session)

	result := s.db.
		Preload("Orders").
		Where(&model.Session{TableID: tableID.String()}).
		First(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	session, err := s.mapper.FromModel(sessionModel)
	if err != nil {
		return nil, err
	}

	return session, err
}

type SessionView struct {
	db *gorm.DB
}

func NewSessionView(db *gorm.DB) *SessionView {
	return &SessionView{
		db: db,
	}
}

func (sv *SessionView) OrderCreatedEventView(sessionID, orderID uuid.UUID) (*dto.OrderCreatedEventView, error) {
	sessionModel := new(model.Session)
	orderModel := new(model.SessionOrder)

	result := sv.db.
		Preload("Table").
		Where(&model.Session{ID: sessionID.String()}).
		First(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	result = sv.db.
		Preload("Customer").
		Preload("MenuItem").
		Where(&model.SessionOrder{
			ID:        orderID.String(),
			SessionID: sessionID.String(),
		}).
		First(orderModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.OrderCreatedEventView{
		RestaurantID: sessionModel.Table.RestaurantID,
		Table:        sessionModel.Table.Name,
		OrderedBy:    orderModel.Customer.FullName,
		MenuItem:     orderModel.MenuItem.Name,
		Quantity:     orderModel.Quantity,
	}, nil
}

func (sv *SessionView) SessionClosedEventView(sessionID uuid.UUID) (*dto.SessionClosedEventView, error) {
	sessionModel := new(model.Session)

	result := sv.db.
		Unscoped().
		Preload("Table").
		Where(&model.Session{ID: sessionID.String()}).
		First(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.SessionClosedEventView{
		RestaurantID: sessionModel.Table.RestaurantID,
		Table:        sessionModel.Table.Name,
	}, nil
}

func (sv *SessionView) OrderCustomerView(customerID, tableID uuid.UUID) ([]dto.OrderCustomerView, error) {
	sessionModel := new(model.Session)
	var orderModels []model.SessionOrder

	result := sv.db.
		Preload("Table").
		Where(&model.Session{TableID: tableID.String()}).
		First(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	result = sv.db.
		Preload(clause.Associations).
		Where(&model.SessionOrder{SessionID: sessionModel.ID}).
		Where(&model.SessionOrder{OrderedBy: customerID.String()}).
		Find(&orderModels)

	if result.Error != nil {
		return nil, result.Error
	}

	orderViews := make([]dto.OrderCustomerView, 0)
	for _, order := range orderModels {
		orderViews = append(orderViews, dto.OrderCustomerView{
			MenuItemName: order.MenuItem.Name,
			Quantity:     order.Quantity,
		})
	}

	return orderViews, nil
}

func (sv *SessionView) OrderTableView(tableID uuid.UUID) ([]dto.OrderTableView, error) {
	var orders []dto.OrderTableView

	result := sv.db.
		Table("session_orders").
		Select("menu_items.name AS menu_item_name, SUM(session_orders.quantity) AS quantity").
		Joins("JOIN menu_items ON menu_items.id = session_orders.menu_item_id").
		Joins("JOIN sessions ON sessions.id = session_orders.session_id").
		Where("sessions.table_id = ?", tableID.String()).
		Group("menu_items.name").
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
