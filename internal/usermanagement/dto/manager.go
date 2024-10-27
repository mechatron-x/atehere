package dto

type Manager struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type ManagerProfile struct {
	Manager Manager `json:"manager,omitempty"`
}
