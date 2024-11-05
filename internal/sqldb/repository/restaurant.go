package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
)

type Restaurant struct {
	queries *dal.Queries
	mapper  mapper.Restaurant
}

func NewRestaurant(db *sql.DB) *Restaurant {
	return &Restaurant{
		queries: dal.New(db),
		mapper:  mapper.NewRestaurant(),
	}
}

func (r *Restaurant) Save(restaurant *aggregate.Restaurant) error {
	restaurantModel := r.mapper.FromAggregate(restaurant)
	saveParams := dal.SaveRestaurantParams(restaurantModel)

	err := r.queries.SaveRestaurant(context.Background(), saveParams)
	if err != nil {
		return r.wrapError(err)
	}

	return nil
}

func (r *Restaurant) wrapError(err error) error {
	return fmt.Errorf("repository.Restaurant: %v", err)
}
