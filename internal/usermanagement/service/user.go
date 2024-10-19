package service

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type User struct {
	userRepository     port.UserRepository
	authInfrastructure port.AuthInfrastructure
}

const (
	model = "user"
)

func (us *User) SignUp(email, password, fullName, birthDate string) (*aggregate.User, error) {
	name, err := valueobject.NewFullName(fullName)
	if err != nil {
		return nil, core.ErrModelValidation(model, err)
	}

	date, err := valueobject.NewBirthDate(birthDate)
	if err != nil {
		return nil, core.ErrModelValidation(model, err)
	}

	user, err := aggregate.NewUser(core.NewAggregate(), name, date)
	if err != nil {
		return nil, core.ErrModelValidation(model, err)
	}

	err = us.authInfrastructure.CreateUser(email, password, user)
	if err != nil {
		return nil, core.ErrModelCreation(model, err)
	}

	err = us.userRepository.Save(user)
	if err != nil {
		return nil, core.ErrModelPersistence(model, err)
	}

	return user, nil
}

func (us *User) GetProfile(idToken string) (*aggregate.User, error) {
	authUser, err := us.authInfrastructure.VerifyUser(idToken)
	if err != nil {
		return nil, err
	}

	user, err := us.userRepository.GetByID(authUser.UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
