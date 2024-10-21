package mapper

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/sqldb/query"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) FromModel(model *query.User) (*aggregate.User, error) {
	id, err := uuid.Parse(model.ID)
	if err != nil {
		return nil, err
	}

	root := core.LoadAggregate(
		id,
		model.CreatedAt.Time,
		model.UpdatedAt.Time,
		&model.DeletedAt.Time)

	fullName, err := valueobject.NewFullName(model.FullName)
	if err != nil {
		return nil, err
	}

	birthDate, err := valueobject.NewBirthDate(model.BirthDate.Time.Format(valueobject.BirthDateLayout))
	if err != nil {
		return nil, err
	}

	return aggregate.LoadUser(root, fullName, birthDate), nil
}

func (u User) FromAggregate(user *aggregate.User) *query.User {
	return &query.User{
		ID:       user.ID().String(),
		FullName: user.FullName().String(),
		BirthDate: sql.NullTime{
			Time:  user.BirthDate().Date(),
			Valid: true,
		},
		CreatedAt: sql.NullTime{
			Time:  user.CreatedAt(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  user.UpdatedAt(),
			Valid: true,
		},
		DeletedAt: sql.NullTime{
			Time:  user.DeletedAt(),
			Valid: user.IsDeleted(),
		},
	}
}
