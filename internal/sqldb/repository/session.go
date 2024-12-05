package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/model"
	"gorm.io/gorm"
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

func (sv *SessionView) OrderCreatedEventView(sessionID, orderID uuid.UUID) (*dto.OrderCreatedEvent, error) {
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

	return &dto.OrderCreatedEvent{
		RestaurantID: sessionModel.Table.RestaurantID,
		Table:        sessionModel.Table.Name,
		OrderedBy:    orderModel.Customer.FullName,
		MenuItem:     orderModel.MenuItem.Name,
		Quantity:     orderModel.Quantity,
	}, nil
}

func (sv *SessionView) SessionClosedEventView(sessionID uuid.UUID) (*dto.SessionClosedEvent, error) {
	sessionModel := new(model.Session)

	result := sv.db.
		Unscoped().
		Preload("Table").
		Where(&model.Session{ID: sessionID.String()}).
		First(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.SessionClosedEvent{
		RestaurantID: sessionModel.Table.RestaurantID,
		Table:        sessionModel.Table.Name,
	}, nil
}

func (sv *SessionView) CustomerOrders(customerID, tableID uuid.UUID) ([]dto.Order, error) {
	orders := make([]dto.Order, 0)
	result := sv.db.
		Table("session_orders").
		Select(`
			menu_items.id AS menu_item_id,
			menu_items.price_amount, 
			menu_items.price_currency,
			menu_items.name AS menu_item_name, 
			SUM(session_orders.quantity) AS quantity
		`).
		Joins("JOIN menu_items ON menu_items.id = session_orders.menu_item_id").
		Joins("JOIN sessions ON sessions.id = session_orders.session_id").
		Where("sessions.table_id = ? AND session_orders.ordered_by = ?",
			tableID.String(),
			customerID.String(),
		).
		Group(`
			menu_items.id,
			menu_items.name,
			menu_items.price_amount,
			menu_items.price_currency
		`).
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (sv *SessionView) TableOrders(tableID uuid.UUID) ([]dto.Order, error) {
	orders := make([]dto.Order, 0)

	result := sv.db.
		Table("session_orders").
		Select(`
			menu_items.id AS menu_item_id,
			menu_items.price_amount, 
			menu_items.price_currency,
			menu_items.name AS menu_item_name, 
			SUM(session_orders.quantity) AS quantity
		`).
		Joins("JOIN menu_items ON menu_items.id = session_orders.menu_item_id").
		Joins("JOIN sessions ON sessions.id = session_orders.session_id").
		Where("sessions.table_id = ?", tableID.String()).
		Group(`
			menu_items.id,
			menu_items.name,
			menu_items.price_amount,
			menu_items.price_currency
		`).
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
