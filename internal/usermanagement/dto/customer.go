package dto

type (
	Customer struct {
		Email     string `json:"email"`
		FullName  string `json:"full_name"`
		Gender    string `json:"gender,omitempty"`
		BirthDate string `json:"birth_date"`
	}

	CustomerSignUp struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FullName  string `json:"full_name"`
		Gender    string `json:"gender,omitempty"`
		BirthDate string `json:"birth_date"`
	}
)
