package service

import (
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/usermanagement/port"
)

type User struct {
	userRepository     port.UserRepository
	authInfrastructure port.AuthInfrastructure
}

func (us *User) SignUp(fullName string, birthDate string) (*aggregate.User, error) {
	name, err := valueobject.NewFullName(fullName)
	if err != nil {
		return nil, err
	}

	date, err := valueobject.NewBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	user := aggregate.NewUser(name, date)

	err = us.authInfrastructure.CreateUser(user)
	if err != nil {
		return nil, err
	}

	user, err = us.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *User) GetUser(idToken string) (*aggregate.User, error) {
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
