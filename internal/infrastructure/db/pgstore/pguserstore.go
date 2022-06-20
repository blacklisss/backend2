package pgstore

import (
	"backend2/internal/entities/userentity"
	"backend2/internal/usecases/app/repos/userrepo"
	"context"
	"database/sql"
	"time"
)

var _ userrepo.UserStore = &Users{}

type User struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
}

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {

	un := &Users{
		db: db,
	}
	return un
}

func (u *Users) CreateUser(ctx context.Context, user *userrepo.User) (*userentity.User, error) {
	us := &userentity.User{}
	return us, nil
}

func (u *Users) FindUser(ctx context.Context, id int64) (*userentity.User, error) {
	us := &userentity.User{}
	return us, nil
}

func (u *Users) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
