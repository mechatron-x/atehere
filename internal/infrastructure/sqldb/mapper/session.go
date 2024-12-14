package mapper

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/infrastructure/sqldb/model"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/session/domain/entity"
	"github.com/mechatron-x/atehere/internal/session/domain/valueobject"
	"gorm.io/gorm"
)

type Session struct{}

func (s Session) FromModel(model *model.Session) (*aggregate.Session, error) {
	session := aggregate.NewSession()

	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	tableID, err := uuid.Parse(model.TableID)
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Order, 0)
	for _, o := range model.Orders {
		orderID, err := uuid.Parse(o.ID)
		if err != nil {
			return nil, err
		}

		menuItemID, err := uuid.Parse(o.MenuItemID)
		if err != nil {
			return nil, err
		}

		orderedBy, err := uuid.Parse(o.OrderedBy)
		if err != nil {
			return nil, err
		}

		quantity, err := valueobject.NewQuantity(o.Quantity)
		if err != nil {
			return nil, err
		}

		order := entity.NewOrder()
		order.SetID(orderID)
		order.SetMenuItemID(menuItemID)
		order.SetOrderedBy(orderedBy)
		order.SetQuantity(quantity)
		order.SetCreatedAt(o.CreatedAt)
		order.SetUpdatedAt(o.UpdatedAt)

		orders = append(orders, order)
	}

	session.SetID(id)
	session.SetTableID(tableID)
	session.SetStartTime(model.StartTime)
	session.SetEndTime(model.EndTime)
	session.SetOrders(orders)
	session.SetCreatedAt(model.CreatedAt)
	session.SetUpdatedAt(model.UpdatedAt)
	if model.DeletedAt.Valid {
		session.SetDeletedAt(model.DeletedAt.Time)
	}

	return session, nil
}

func (s Session) FromModels(models []model.Session) ([]*aggregate.Session, error) {
	aggregates := make([]*aggregate.Session, 0)

	for _, m := range models {
		aggregate, err := s.FromModel(&m)
		if err != nil {
			return nil, err
		}

		aggregates = append(aggregates, aggregate)
	}

	return aggregates, nil
}

func (s Session) FromAggregate(aggregate *aggregate.Session) *model.Session {
	orders := make([]model.SessionOrder, 0)
	for _, o := range aggregate.Orders() {
		order := model.SessionOrder{
			ID:         o.ID().String(),
			SessionID:  aggregate.ID().String(),
			MenuItemID: o.MenuItemID().String(),
			OrderedBy:  o.OrderedBy().String(),
			Quantity:   o.Quantity().Int(),
			CreatedAt:  o.CreatedAt(),
			UpdatedAt:  o.UpdatedAt(),
		}

		orders = append(orders, order)
	}

	return &model.Session{
		ID:        aggregate.ID().String(),
		TableID:   aggregate.TableID().String(),
		StartTime: aggregate.StartTime(),
		EndTime:   aggregate.EndTime(),
		Orders:    orders,
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

func (s Session) FromAggregates(aggregates []*aggregate.Session) []*model.Session {
	models := make([]*model.Session, 0)
	for _, a := range aggregates {
		models = append(models, s.FromAggregate(a))
	}

	return models
}
