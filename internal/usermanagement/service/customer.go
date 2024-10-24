package service

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type Customer struct {
	customerRepo port.CustomerRepository
	customerAuth port.CustomerAuthenticator
}

const (
	context = "service.Customer"
)

func NewCustomer(userRepository port.CustomerRepository, authInfrastructure port.CustomerAuthenticator) *Customer {
	return &Customer{
		customerRepo: userRepository,
		customerAuth: authInfrastructure,
	}
}

func (cs *Customer) SignUp(customerDto dto.Customer) (*dto.SignUpCustomer, core.DomainError) {
	customerSignUpDto := &dto.SignUpCustomer{
		DateFormat: valueobject.BirthDateLayoutANSIC,
		Genders:    valueobject.GetGenders(),
	}

	customer, err := cs.validateSignUpDto(customerDto)
	if err != nil {
		return nil, core.NewValidationError(context, err)
	}

	// err = cs.customerAuth.CreateUser(customer)
	// if err != nil {
	// 	return nil, core.ErrModelCreation(model, err)
	// }

	err = cs.customerRepo.Save(customer)
	if err != nil {
		return nil, core.NewPersistenceError(context, err)
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
