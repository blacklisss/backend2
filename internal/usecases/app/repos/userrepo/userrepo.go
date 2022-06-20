package userrepo

import (
	"backend2/internal/entities/userentity"
	"context"
)

type User struct {
	Name string
}

type UserStore interface {
	CreateUser(ctx context.Context, user *User) (*userentity.User, error)
	FindUser(ctx context.Context, id int64) (*userentity.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore: ustore,
	}
}
