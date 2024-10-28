package dto

type (
	Manager struct {
		Email       string `json:"email,omitempty"`
		FullName    string `json:"full_name,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
	}

	ManagerSignUp struct {
		Email       string `json:"email,omitempty"`
		Password    string `json:"password,omitempty"`
		FullName    string `json:"full_name,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
	}
)
