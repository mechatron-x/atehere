package response

import "github.com/mechatron-x/atehere/internal/restaurant/dto"

type (
	RestaurantCreate struct {
		*dto.Restaurant
	}

	RestaurantList struct {
		Restaurants []dto.Restaurant `json:"restaurants"`
	}
)
