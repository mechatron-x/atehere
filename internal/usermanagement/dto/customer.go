package dto

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

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

func (c Customer) Update(customer *aggregate.Customer) error {
	if !core.IsEmptyString(c.FullName) {
		verifiedFullName, err := valueobject.NewFullName(c.FullName)
		if err != nil {
			return err
		}
		customer.SetFullName(verifiedFullName)
	}

	if !core.IsEmptyString(c.BirthDate) {
		verifiedBirthDate, err := valueobject.NewBirthDate(c.BirthDate)
		if err != nil {
			return err
		}
		customer.SetBirthDate(verifiedBirthDate)
	}

	if !core.IsEmptyString(c.Gender) {
		verifiedGender := valueobject.ParseGender(c.Gender)
		customer.SetGender(verifiedGender)
	}

	return nil
}

func (cs CustomerSignUp) Validate() (*aggregate.Customer, error) {
	verifiedEmail, err := valueobject.NewEmail(cs.Email)
	if err != nil {
		return nil, err
	}

	verifiedPassword, err := valueobject.NewPassword(cs.Password)
	if err != nil {
		return nil, err
	}

	verifiedFullName, err := valueobject.NewFullName(cs.FullName)
	if err != nil {
		return nil, err
	}

	verifiedGender := valueobject.ParseGender(cs.Gender)

	verifiedBirthDate, err := valueobject.NewBirthDate(cs.BirthDate)
	if err != nil {
		return nil, err
	}

	customer := aggregate.NewCustomer()
	customer.SetEmail(verifiedEmail)
	customer.SetPassword(verifiedPassword)
	customer.SetFullName(verifiedFullName)
	customer.SetGender(verifiedGender)
	customer.SetBirthDate(verifiedBirthDate)

	return customer, nil
}
