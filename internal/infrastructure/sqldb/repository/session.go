package repository

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db     *gorm.DB
	mapper mapper.Session
}

func NewSession(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db:     db,
		mapper: mapper.Session{},
	}
}

func (s *SessionRepository) Save(session *aggregate.Session) error {
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

func (s *SessionRepository) GetByTableID(tableID uuid.UUID) (*aggregate.Session, error) {
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

func (s *SessionRepository) GetByID(ID uuid.UUID) (*aggregate.Session, error) {
	sessionModel := new(model.Session)

	result := s.db.
		Unscoped().
		Preload("Orders").
		Where(&model.Session{ID: ID.String()}).
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

func (sv *SessionView) OrderCreatedEventView(sessionID, orderID uuid.UUID) (*dto.NewOrderEvent, error) {
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

	return &dto.NewOrderEvent{
		RestaurantID: sessionModel.Table.RestaurantID,
		Table:        sessionModel.Table.Name,
		OrderedBy:    orderModel.Customer.FullName,
		MenuItem:     orderModel.MenuItem.Name,
		Quantity:     orderModel.Quantity,
	}, nil
}

func (sv *SessionView) CheckoutEventView(sessionID uuid.UUID) (*dto.CheckoutEvent, error) {
	sessionModel := new(model.Session)

	result := sv.db.
		Unscoped().
		Preload("Table").
		Where(&model.Session{ID: sessionID.String()}).
		First(sessionModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.CheckoutEvent{
		RestaurantID: sessionModel.Table.RestaurantID,
		Table:        sessionModel.Table.Name,
	}, nil
}

func (sv *SessionView) GetTableOrdersView(sessionID uuid.UUID) ([]dto.TableOrderView, error) {
	var orders []dto.TableOrderView

	result := sv.db.Table("table_orders").
		Where("session_id = ?", sessionID.String()).
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (sv *SessionView) GetManagerOrdersView(sessionID uuid.UUID) ([]dto.ManagerOrderView, error) {
	var orders []dto.ManagerOrderView

	result := sv.db.Table("manager_orders").
		Where("session_id = ?", sessionID.String()).
		Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
