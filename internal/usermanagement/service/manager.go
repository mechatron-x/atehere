package service

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type Manager struct {
	managerRepo   port.ManagerRepository
	authenticator port.Authenticator
}

func NewManager(
	managerRepository port.ManagerRepository,
	authInfrastructure port.Authenticator,
) *Manager {
	return &Manager{
		managerRepo:   managerRepository,
		authenticator: authInfrastructure,
	}
}

func (ms *Manager) SignUp(signUpDto dto.ManagerSignUp) (*dto.Manager, error) {
	manager, err := ms.validateSignUpDto(signUpDto)
	if err != nil {
		return nil, core.ErrValidation
	}

	portErr := ms.authenticator.CreateUser(
		manager.ID().String(),
		manager.Email().String(),
		manager.Password().String(),
	)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	savedManager, portErr := ms.managerRepo.Save(manager)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	savedManager.SetEmail(manager.Email())
	manager.SetPhoneNumber(manager.PhoneNumber())

	return &dto.Manager{
		Email:       savedManager.Email().String(),
		FullName:    savedManager.FullName().String(),
		PhoneNumber: savedManager.PhoneNumber().String(),
	}, nil
}

func (ms *Manager) GetProfile(idToken string) (*dto.Manager, error) {
	uid, portErr := ms.authenticator.GetUserID(idToken)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	email, portErr := ms.authenticator.GetUserEmail(idToken)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	manager, portErr := ms.managerRepo.GetByID(uid)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	verifiedEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, core.ErrValidation
	}
	manager.SetEmail(verifiedEmail)

	return &dto.Manager{
		Email:       manager.Email().String(),
		FullName:    manager.FullName().String(),
		PhoneNumber: manager.PhoneNumber().String(),
	}, nil
}

func (ms *Manager) UpdateProfile(idToken string, updateDto dto.Manager) (*dto.Manager, error) {
	uid, portErr := ms.authenticator.GetUserID(idToken)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	email, portErr := ms.authenticator.GetUserEmail(idToken)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	manager, portErr := ms.managerRepo.GetByID(uid)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	err := ms.updateManager(updateDto, manager)
	if err != nil {
		return nil, core.ErrValidation
	}

	manager, portErr = ms.managerRepo.Save(manager)
	if portErr != nil {
		return nil, core.MapPortErrorToDomain(portErr)
	}

	verifiedEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, core.ErrValidation
	}
	manager.SetEmail(verifiedEmail)

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
