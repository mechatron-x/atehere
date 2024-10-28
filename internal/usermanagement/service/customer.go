package service

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type Customer struct {
	customerRepo  port.CustomerRepository
	authenticator port.Authenticator
}

func NewCustomer(
	customerRepository port.CustomerRepository,
	authInfrastructure port.Authenticator,
) *Customer {
	return &Customer{
		customerRepo:  customerRepository,
		authenticator: authInfrastructure,
	}
}

func (cs *Customer) SignUp(customerDto dto.CustomerSignUp) (*dto.Customer, error) {
	customer, err := cs.validateSignUpDto(customerDto)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	err = cs.authenticator.CreateUser(
		customer.ID().String(),
		customer.Email().String(),
		customer.Password().String(),
	)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	err = cs.customerRepo.Save(customer)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return &dto.Customer{
		Email:     customer.Email().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().String(),
	}, nil
}

func (cs *Customer) GetProfile(idToken string) (*dto.Customer, error) {
	customer, err := cs.getCustomer(idToken)
	if err != nil {
		return nil, err
	}

	return &dto.Customer{
		Email:     customer.Email().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().String(),
	}, nil
}

func (cs *Customer) UpdateProfile(idToken string, customerDto dto.Customer) (*dto.Customer, error) {
	customer, err := cs.getCustomer(idToken)
	if err != nil {
		return nil, err
	}

	err = cs.updateCustomer(customerDto, customer)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	err = cs.customerRepo.Save(customer)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return &dto.Customer{
		Email:     customer.Email().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().String(),
	}, nil
}

func (cs *Customer) updateCustomer(customerDto dto.Customer, customer *aggregate.Customer) error {
	if !core.IsEmptyString(customerDto.FullName) {
		verifiedFullName, err := valueobject.NewFullName(customerDto.FullName)
		if err != nil {
			return err
		}
		customer.SetFullName(verifiedFullName)
	}

	if !core.IsEmptyString(customerDto.BirthDate) {
		verifiedBirthDate, err := valueobject.NewBirthDate(customerDto.BirthDate)
		if err != nil {
			return err
		}
		customer.SetBirthDate(verifiedBirthDate)
	}

	if !core.IsEmptyString(customerDto.Gender) {
		verifiedGender := valueobject.ParseGender(customerDto.Gender)
		customer.SetGender(verifiedGender)
	}

	return nil
}

func (cs *Customer) validateSignUpDto(signUpDto dto.CustomerSignUp) (*aggregate.Customer, error) {
	verifiedEmail, err := valueobject.NewEmail(signUpDto.Email)
	if err != nil {
		return nil, err
	}

	verifiedPassword, err := valueobject.NewPassword(signUpDto.Password)
	if err != nil {
		return nil, err
	}

	verifiedFullName, err := valueobject.NewFullName(signUpDto.FullName)
	if err != nil {
		return nil, err
	}

	verifiedGender := valueobject.ParseGender(signUpDto.Gender)

	verifiedBirthDate, err := valueobject.NewBirthDate(signUpDto.BirthDate)
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

func (cs *Customer) getCustomer(idToken string) (*aggregate.Customer, error) {
	id, err := cs.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	email, err := cs.authenticator.GetUserEmail(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	customer, err := cs.customerRepo.GetByID(id)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	verifiedEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}
	customer.SetEmail(verifiedEmail)

	return customer, nil
}
