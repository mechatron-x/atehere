package request

type (
	SignUpUser struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FullName  string `json:"full_name"`
		BirthDate string `json:"birth_date"`
	}
)
