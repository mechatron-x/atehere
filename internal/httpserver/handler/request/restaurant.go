package request

import "github.com/mechatron-x/atehere/internal/restaurant/dto"

type (
	RestaurantCreate struct {
		dto.RestaurantCreate
	}

	RestaurantFilter struct {
		Page int `json:"page"`
		dto.RestaurantFilter
	}
)
