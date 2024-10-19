package port

import "github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"

type UserRepository interface {
	Save(user *aggregate.User) error
	GetByID(id string) (*aggregate.User, error)
}
