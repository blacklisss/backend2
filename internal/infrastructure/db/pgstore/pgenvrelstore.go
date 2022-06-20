package pgstore

import (
	"backend2/internal/entities/enventity"
	"backend2/internal/entities/userentity"
	"backend2/internal/usecases/app/repos/envrepo"
	"context"
	"database/sql"
)

var _ envrepo.EnvRelationStore = &EnvRels{}

type EnvRels struct {
	db *sql.DB
}

func NewEnvRels(db *sql.DB) *EnvRels {
	enr := &EnvRels{
		db: db,
	}
	return enr
}

func (er *EnvRels) AddUserToGroup(ctx context.Context, userid int64, envid int64) error {
	return nil
}

func (er *EnvRels) RemoveUserFromGroup(ctx context.Context, userid int64, envid int64) error {
	return nil
}

func (er *EnvRels) FindUserGroups(ctx context.Context, name string) ([]*enventity.UserEnviroment, error) {
	return nil, nil
}

func (er *EnvRels) FindUsersByGroup(ctx context.Context, name string) ([]*userentity.User, error) {
	return nil, nil
}
