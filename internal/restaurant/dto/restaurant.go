package dto

type (
	RestaurantCreate struct {
		Name           string   `json:"name"`
		FoundationYear string   `json:"foundation_year"`
		PhoneNumber    string   `json:"phone_number"`
		OpeningTime    string   `json:"opening_time"`
		ClosingTime    string   `json:"closing_time"`
		WorkingDays    []string `json:"working_days"`
	}

	Restaurant struct {
		ID             string   `json:"id"`
		Owner          Owner    `json:"owner"`
		Name           string   `json:"name"`
		FoundationYear string   `json:"foundation_year"`
		PhoneNumber    string   `json:"phone_number"`
		OpeningTime    string   `json:"opening_time"`
		ClosingTime    string   `json:"closing_time"`
		WorkingDays    []string `json:"working_days"`
	}

	Owner struct {
		ID       string `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}
)
