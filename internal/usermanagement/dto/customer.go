package dto

type (
	Customer struct {
		Email     string `json:"email,omitempty"`
		FullName  string `json:"full_name,omitempty"`
		Gender    string `json:"gender,omitempty"`
		BirthDate string `json:"birth_date,omitempty"`
	}

	CustomerSignUp struct {
		Email     string `json:"email,omitempty"`
		Password  string `json:"password,omitempty"`
		FullName  string `json:"full_name,omitempty"`
		Gender    string `json:"gender,omitempty"`
		BirthDate string `json:"birth_date,omitempty"`
	}
)
