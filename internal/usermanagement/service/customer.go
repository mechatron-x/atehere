package service

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type Customer struct {
	userRepository     port.CustomerRepository
	authInfrastructure port.CustomerAuthenticator
}

const (
	model = "customer"
)

func NewCustomer(userRepository port.CustomerRepository, authInfrastructure port.CustomerAuthenticator) *Customer {
	return &Customer{
		userRepository:     userRepository,
		authInfrastructure: authInfrastructure,
	}
}

func (cs *Customer) SignUp(customerDto dto.Customer) (dto.SignUpCustomer, error) {
	customerSignUpDto := dto.SignUpCustomer{
		DateFormat: valueobject.BirthDateLayoutANSIC,
		Genders:    valueobject.GetGenders(),
	}

	customer, err := cs.validateSignUpDto(customerDto)
	if err != nil {
		return customerSignUpDto, core.ErrModelValidation(model, err)
	}

	// err = cs.authInfrastructure.CreateUser(user)
	// if err != nil {
	// 	return nil, core.ErrModelCreation(model, err)
	// }

	err = cs.userRepository.Save(customer)
	if err != nil {
		return customerSignUpDto, core.ErrModelPersistence(model, err)
	}

	customerSignUpDto.Customer = dto.Customer{
		ID:        customer.ID().String(),
		Email:     customer.Email().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().String(),
	}

	return customerSignUpDto, nil
}

func (cs *Customer) validateSignUpDto(signUpDto dto.Customer) (*aggregate.Customer, error) {
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
