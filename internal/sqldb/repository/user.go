package repository

import (
	"context"
	"database/sql"

	"github.com/mechatron-x/atehere/internal/sqldb/mapper"
	"github.com/mechatron-x/atehere/internal/sqldb/query"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/aggregate"
)

type User struct {
	q *query.Queries
	m mapper.User
}

func NewUser(db *sql.DB) *User {
	q := query.New(db)
	m := mapper.NewUser()
	return &User{
		q: q,
		m: m,
	}
}

func (u *User) Save(user *aggregate.User) error {
	userModel := u.m.FromAggregate(user)
	saveParams := query.SaveUserParams{
		ID:        userModel.ID,
		FullName:  userModel.FullName,
		BirthDate: userModel.BirthDate,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		DeletedAt: userModel.DeletedAt,
	}

	_, err := u.q.SaveUser(context.Background(), saveParams)
	return err
}

func (u *User) GetByID(id string) (*aggregate.User, error) {
	userModel, err := u.q.GetUserByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	user, err := u.m.FromModel(&userModel)
	if err != nil {
		return nil, err
	}

	return user, nil
}
