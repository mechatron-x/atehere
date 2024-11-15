package service

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type Customer struct {
	customerRepo port.CustomerRepository
	authService  port.Authenticator
}

func NewCustomer(
	customerRepository port.CustomerRepository,
	authService port.Authenticator,
) *Customer {
	return &Customer{
		customerRepo: customerRepository,
		authService:  authService,
	}
}

func (cs *Customer) SignUp(customerDto *dto.CustomerSignUp) (*dto.Customer, error) {
	customer, err := customerDto.Validate()
	if err != nil {
		logger.Error("Cannot map customer dto to aggregate", err)
		return nil, core.NewValidationFailureError(err)
	}

	err = cs.authService.CreateUser(
		customer.ID().String(),
		customer.Email().String(),
		customer.Password().String(),
	)
	if err != nil {
		logger.Error("Failed to authenticate customer", err)
		return nil, core.NewPersistenceFailureError(err)
	}

	err = cs.customerRepo.Save(customer)
	if err != nil {
		logger.Error("Failed to save customer to db", err)
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
		logger.Error("Cannot get customer with id token", err)
		return nil, err
	}

	return &dto.Customer{
		Email:     customer.Email().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().String(),
	}, nil
}

func (cs *Customer) UpdateProfile(idToken string, customerDto *dto.Customer) (*dto.Customer, error) {
	customer, err := cs.getCustomer(idToken)
	if err != nil {
		logger.Error("Cannot get customer with id token", err)
		return nil, err
	}

	err = customerDto.Update(customer)
	if err != nil {
		logger.Error("Cannot map customer update dto to aggregate", err)
		return nil, core.NewValidationFailureError(err)
	}

	err = cs.customerRepo.Save(customer)
	if err != nil {
		logger.Error("Failed to save customer to db", err)
		return nil, core.NewPersistenceFailureError(err)
	}

	return &dto.Customer{
		Email:     customer.Email().String(),
		FullName:  customer.FullName().String(),
		Gender:    customer.Gender().String(),
		BirthDate: customer.BirthDate().String(),
	}, nil
}

func (cs *Customer) getCustomer(idToken string) (*aggregate.Customer, error) {
	id, err := cs.authService.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	email, err := cs.authService.GetUserEmail(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	customer, err := cs.customerRepo.GetByID(uuid.MustParse(id))
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
