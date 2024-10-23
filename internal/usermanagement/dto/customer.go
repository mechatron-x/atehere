package dto

type CustomerSignUp struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}
