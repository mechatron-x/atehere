package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/sqldb/dal"
	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
)

const (
	DefaultPageSize int = 10
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

func (r *Restaurant) GetAll(page int) ([]*aggregate.Restaurant, error) {
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
		restaurant, err := r.mapper.FromModel(model)
		if err != nil {
			return nil, r.wrapError(err)
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func (r *Restaurant) wrapError(err error) error {
	return fmt.Errorf("repository.Restaurant: %v", err)
}
