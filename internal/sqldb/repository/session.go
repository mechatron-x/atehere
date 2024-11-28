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
	var sessionModel model.Session

	result := s.db.Model(&model.Session{TableID: tableID.String()}).
		Preload("Orders").
		First(&sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	session, err := s.mapper.FromModel(&sessionModel)
	if err != nil {
		return nil, err
	}

	return session, err
}

func (s *Session) HasActiveSessions(tableID uuid.UUID) bool {
	var count int64
	result := s.db.Model(&model.Session{TableID: tableID.String()}).
		Count(&count)

	if result.Error != nil {
		return false
	}

	return count != 0
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
	var session model.Session
	var order model.SessionOrder

	result := sv.db.Model(&model.Session{ID: sessionID.String()}).
		Preload("Table").
		First(&session)
	if result.Error != nil {
		return nil, result.Error
	}

	result = sv.db.Model(&model.SessionOrder{ID: orderID.String()}).
		Preload(clause.Associations).
		First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.OrderCreatedEventView{
		RestaurantID: session.Table.RestaurantID,
		Table:        session.Table.Name,
		OrderedBy:    order.Customer.FullName,
		MenuItem:     order.MenuItem.Name,
		Quantity:     order.Quantity,
	}, nil
}

func (sv *SessionView) OrderCustomerView(customerID uuid.UUID) ([]dto.OrderCustomerView, error) {
	var orders []model.SessionOrder

	result := sv.db.Model(&model.SessionOrder{OrderedBy: customerID.String()}).
		Preload(clause.Associations).
		Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	orderViews := make([]dto.OrderCustomerView, 0)
	for _, order := range orders {
		orderViews = append(orderViews, dto.OrderCustomerView{
			MenuItemName: order.MenuItem.Name,
			Quantity:     order.Quantity,
		})
	}

	return orderViews, nil
}
