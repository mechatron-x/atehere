package response

type (
	SignUpUser struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		FullName  string `json:"full_name,omitempty"`
		BirthDate string `json:"birth_date,omitempty"`
	}
)
