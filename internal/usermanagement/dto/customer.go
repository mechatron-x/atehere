package dto

type Customer struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type SignUpCustomer struct {
	DateFormat string   `json:"date_format,omitempty"`
	Genders    []string `json:"genders,omitempty"`
	Customer   Customer `json:"customer,omitempty"`
}

type CustomerProfile struct {
	DateFormat string   `json:"date_format,omitempty"`
	Genders    []string `json:"genders,omitempty"`
	Customer   Customer `json:"customer,omitempty"`
}
