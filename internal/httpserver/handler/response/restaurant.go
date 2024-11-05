package response

import "github.com/mechatron-x/atehere/internal/restaurant/dto"

type (
	RestaurantCreate struct {
		*dto.Restaurant
	}

	RestaurantList struct {
		AvailableWorkingDays []string         `json:"available_working_days"`
		FoundationYearFormat string           `json:"foundation_year_format"`
		WorkingTimeFormat    string           `json:"working_time_format"`
		Restaurants          []dto.Restaurant `json:"restaurants"`
	}
)
