package dto

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

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

func (m Manager) Update(manager *aggregate.Manager) error {
	if !core.IsEmptyString(m.FullName) {
		verifiedFullName, err := valueobject.NewFullName(m.FullName)
		if err != nil {
			return err
		}
		manager.SetFullName(verifiedFullName)
	}

	if !core.IsEmptyString(m.PhoneNumber) {
		verifiedPhoneNumber, err := valueobject.NewPhoneNumber(m.PhoneNumber)
		if err != nil {
			return err
		}
		manager.SetPhoneNumber(verifiedPhoneNumber)
	}

	return nil
}

func (ms ManagerSignUp) Validate() (*aggregate.Manager, error) {
	verifiedEmail, err := valueobject.NewEmail(ms.Email)
	if err != nil {
		return nil, err
	}

	verifiedPassword, err := valueobject.NewPassword(ms.Password)
	if err != nil {
		return nil, err
	}

	verifiedFullName, err := valueobject.NewFullName(ms.FullName)
	if err != nil {
		return nil, err
	}

	verifiedPhoneNumber, err := valueobject.NewPhoneNumber(ms.PhoneNumber)
	if err != nil {
		return nil, err
	}

	manager := aggregate.NewManager()
	manager.SetEmail(verifiedEmail)
	manager.SetPassword(verifiedPassword)
	manager.SetFullName(verifiedFullName)
	manager.SetPhoneNumber(verifiedPhoneNumber)

	return manager, nil
}
