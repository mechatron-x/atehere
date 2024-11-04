package service

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/logger"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type Manager struct {
	managerRepo port.ManagerRepository
	authService port.Authenticator
}

func NewManager(
	managerRepository port.ManagerRepository,
	authService port.Authenticator,
) *Manager {
	return &Manager{
		managerRepo: managerRepository,
		authService: authService,
	}
}

func (ms *Manager) SignUp(signUpDto dto.ManagerSignUp) (*dto.Manager, error) {
	manager, err := ms.validateSignUpDto(signUpDto)
	if err != nil {
		logger.Error("Cannot map manager dto to aggregate", err)
		return nil, core.NewValidationFailureError(err)
	}

	err = ms.authService.CreateUser(
		manager.ID().String(),
		manager.Email().String(),
		manager.Password().String(),
	)
	if err != nil {
		logger.Error("Failed to authenticate manager", err)
		return nil, core.NewPersistenceFailureError(err)
	}

	err = ms.managerRepo.Save(manager)
	if err != nil {
		logger.Error("Failed to save manager to db", err)
		return nil, core.NewPersistenceFailureError(err)
	}

	manager.SetPhoneNumber(manager.PhoneNumber())

	return &dto.Manager{
		Email:       manager.Email().String(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
	}, nil
}

func (ms *Manager) GetProfile(idToken string) (*dto.Manager, error) {
	manager, err := ms.getManager(idToken)
	if err != nil {
		logger.Error("Cannot get manager with id token", err)
		return nil, err
	}

	return &dto.Manager{
		Email:       manager.Email().String(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
	}, nil
}

func (ms *Manager) UpdateProfile(idToken string, updateDto dto.Manager) (*dto.Manager, error) {
	manager, err := ms.getManager(idToken)
	if err != nil {
		logger.Error("Cannot get manager with id token", err)
		return nil, err
	}

	err = ms.updateManager(updateDto, manager)
	if err != nil {
		logger.Error("Cannot map manager update dto to aggregate", err)
		return nil, core.NewValidationFailureError(err)
	}

	err = ms.managerRepo.Save(manager)
	if err != nil {
		logger.Error("Failed to save manager to db", err)
		return nil, core.NewPersistenceFailureError(err)
	}

	return &dto.Manager{
		Email:       manager.Email().String(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
	}, nil
}

func (ms *Manager) updateManager(updateDto dto.Manager, manager *aggregate.Manager) error {
	if !core.IsEmptyString(updateDto.FullName) {
		verifiedFullName, err := valueobject.NewFullName(updateDto.FullName)
		if err != nil {
			return err
		}
		manager.SetFullName(verifiedFullName)
	}

	if !core.IsEmptyString(updateDto.PhoneNumber) {
		verifiedPhoneNumber, err := valueobject.NewPhoneNumber(updateDto.PhoneNumber)
		if err != nil {
			return err
		}
		manager.SetPhoneNumber(verifiedPhoneNumber)
	}

	return nil
}

func (ms *Manager) validateSignUpDto(signUpDto dto.ManagerSignUp) (*aggregate.Manager, error) {
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

	verifiedPhoneNumber, err := valueobject.NewPhoneNumber(signUpDto.PhoneNumber)
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

func (ms *Manager) getManager(idToken string) (*aggregate.Manager, error) {
	id, err := ms.authService.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	email, err := ms.authService.GetUserEmail(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	manager, err := ms.managerRepo.GetByID(id)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	verifiedEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}
	manager.SetEmail(verifiedEmail)

	return manager, err
}
