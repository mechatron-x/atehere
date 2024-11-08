package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
)

const (
	DefaultPageSize int = 10
)

type Restaurant struct {
	db      *sql.DB
	queries *dal.Queries
	rMapper mapper.Restaurant
	tMapper mapper.Table
}

func NewRestaurant(db *sql.DB) *Restaurant {
	return &Restaurant{
		db:      db,
		queries: dal.New(db),
		rMapper: mapper.NewRestaurant(),
		tMapper: mapper.NewTable(),
	}
}

func (r *Restaurant) Save(restaurant *aggregate.Restaurant) error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Commit()

	queries := r.queries.WithTx(tx)

	if err := r.saveRestaurant(queries, restaurant); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.deleteTables(queries, restaurant.ID()); err != nil {
		tx.Rollback()
		return err
	}

	if err := r.saveTables(queries, restaurant.ID(), restaurant.Tables()); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *Restaurant) GetAll(page int) ([]*aggregate.Restaurant, error) {
	if page < 0 {
		page = 0
	} else {
		page -= 1
	}

	getParams := dal.GetRestaurantsParams{
		Limit:  int64(DefaultPageSize),
		Offset: int64(page * DefaultPageSize),
	}

	restaurantModels, err := r.queries.GetRestaurants(context.Background(), getParams)
	if err != nil {
		return nil, r.wrapError(err)
	}

	restaurants := make([]*aggregate.Restaurant, 0)

	for _, model := range restaurantModels {
		restaurant, err := r.rMapper.FromModel(model)
		if err != nil {
			return nil, r.wrapError(err)
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func (r *Restaurant) saveRestaurant(queries *dal.Queries, restaurant *aggregate.Restaurant) error {
	restaurantModel := r.rMapper.FromAggregate(restaurant)
	saveParams := dal.SaveRestaurantParams(restaurantModel)

	return queries.SaveRestaurant(context.Background(), saveParams)
}

func (r *Restaurant) saveTables(queries *dal.Queries, restaurantID uuid.UUID, tables []entity.Table) error {
	models := r.tMapper.FromEntities(restaurantID, tables)
	for _, model := range models {
		err := queries.SaveRestaurantTable(context.Background(), dal.SaveRestaurantTableParams(model))
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Restaurant) deleteTables(queries *dal.Queries, restaurantID uuid.UUID) error {
	return queries.DeleteRestaurantTables(context.Background(), restaurantID)
}

func (r *Restaurant) wrapError(err error) error {
	return fmt.Errorf("repository.Restaurant: %v", err)
}
