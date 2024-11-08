package dto

type (
	RestaurantCreate struct {
		Name           string        `json:"name"`
		FoundationYear string        `json:"foundation_year"`
		PhoneNumber    string        `json:"phone_number"`
		OpeningTime    string        `json:"opening_time"`
		ClosingTime    string        `json:"closing_time"`
		WorkingDays    []string      `json:"working_days"`
		Tables         []TableCreate `json:"tables"`
		Image          string        `json:"image"`
	}

	RestaurantSummary struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		PhoneNumber string   `json:"phone_number"`
		OpeningTime string   `json:"opening_time"`
		ClosingTime string   `json:"closing_time"`
		WorkingDays []string `json:"working_days"`
	}

	Restaurant struct {
		ID             string   `json:"id"`
		OwnerID        string   `json:"owner_id"`
		Name           string   `json:"name"`
		FoundationYear string   `json:"foundation_year"`
		PhoneNumber    string   `json:"phone_number"`
		OpeningTime    string   `json:"opening_time"`
		ClosingTime    string   `json:"closing_time"`
		WorkingDays    []string `json:"working_days"`
		ImageURL       string   `json:"image_url"`
		Tables         []Table  `json:"tables"`
	}
)
