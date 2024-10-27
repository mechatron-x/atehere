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

func (ms *Manager) SignUp(managerDto dto.Manager) (*dto.Manager, error) {
	manager, err := ms.validateSignUpDto(managerDto)
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
		ID:          savedManager.ID().String(),
		Email:       savedManager.Email().String(),
		FullName:    savedManager.FullName().String(),
		PhoneNumber: savedManager.PhoneNumber().String(),
	}, nil
}

func (ms *Manager) GetProfile(idToken string) (*dto.ManagerProfile, error) {
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

	managerProfileDto := &dto.ManagerProfile{
		Manager: dto.Manager{
			ID:          manager.ID().String(),
			Email:       manager.Email().String(),
			FullName:    manager.FullName().String(),
			PhoneNumber: manager.PhoneNumber().String(),
		},
	}

	return managerProfileDto, nil
}

func (ms *Manager) validateSignUpDto(signUpDto dto.Manager) (*aggregate.Manager, error) {
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
