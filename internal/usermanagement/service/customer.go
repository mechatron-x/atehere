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

const (
	context = "service.Customer"
)

func NewCustomer(
	userRepository port.CustomerRepository,
	authInfrastructure port.Authenticator,
) *Customer {
	return &Customer{
		customerRepo:  userRepository,
		authenticator: authInfrastructure,
	}
}

func (cs *Customer) SignUp(customerDto dto.Customer) (*dto.Customer, core.DomainError) {
	customer, err := cs.validateSignUpDto(customerDto)
	if err != nil {
		return nil, core.NewValidationError(context, err)
	}

	err = cs.authenticator.CreateUser(
		customer.ID().String(),
		customer.Email().String(),
		customer.Password().String(),
	)
	if err != nil {
		return nil, core.MapPortErrorToDomain(context, err)
	}

	savedCustomer, err := cs.customerRepo.Save(customer)
	if err != nil {
		return nil, core.MapPortErrorToDomain(context, err)
	}

	savedCustomer.SetEmail(customer.Email())

	return &dto.Customer{
		ID:        savedCustomer.ID().String(),
		Email:     savedCustomer.Email().String(),
		FullName:  savedCustomer.FullName().String(),
		Gender:    savedCustomer.Gender().String(),
		BirthDate: savedCustomer.BirthDate().String(),
	}, nil
}

func (cs *Customer) GetProfile(idToken string) (*dto.CustomerProfile, core.DomainError) {
	uid, err := cs.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.MapPortErrorToDomain(context, err)
	}

	email, err := cs.authenticator.GetUserEmail(idToken)
	if err != nil {
		return nil, core.MapPortErrorToDomain(context, err)
	}

	customer, err := cs.customerRepo.GetByID(uid)
	if err != nil {
		return nil, core.MapPortErrorToDomain(context, err)
	}

	verifiedEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, core.MapPortErrorToDomain(context, err)
	}
	customer.SetEmail(verifiedEmail)

	customerProfileDto := &dto.CustomerProfile{
		DateFormat: valueobject.BirthDateLayoutANSIC,
		Genders:    valueobject.GetGenders(),
		Customer: dto.Customer{
			ID:        customer.ID().String(),
			Email:     customer.Email().String(),
			FullName:  customer.FullName().String(),
			Gender:    customer.Gender().String(),
			BirthDate: customer.BirthDate().String(),
		},
	}

	return customerProfileDto, nil
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
