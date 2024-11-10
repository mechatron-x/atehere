package response

import "github.com/mechatron-x/atehere/internal/restaurant/dto"

type (
	RestaurantCreate struct {
		*dto.Restaurant
	}

	Restaurant[TDTO any] struct {
		AvailableWorkingDays []string `json:"available_working_days"`
		FoundationYearFormat string   `json:"foundation_year_format"`
		WorkingTimeFormat    string   `json:"working_time_format"`
		Restaurant           TDTO     `json:"restaurant"`
	}

	RestaurantList[TDTO any] struct {
		AvailableWorkingDays []string `json:"available_working_days"`
		FoundationYearFormat string   `json:"foundation_year_format"`
		WorkingTimeFormat    string   `json:"working_time_format"`
		Restaurants          []TDTO   `json:"restaurants"`
	}
)
