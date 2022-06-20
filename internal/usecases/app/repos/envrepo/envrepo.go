package envrepo

import (
	"backend2/internal/entities/enventity"
	"backend2/internal/entities/userentity"
	"context"
)

type Enviroment struct {
	EnvType int
	Name    string
}

type EnvStore interface {
	CreateEnv(ctx context.Context, env *Enviroment) (*enventity.UserEnviroment, error)
	FindEnv(ctx context.Context, id int64) (*enventity.UserEnviroment, error)
	DeleteEnv(ctx context.Context, id int64) error
}

type Envs struct {
	estore EnvStore
}

func NewEnvs(estore EnvStore) *Envs {
	return &Envs{
		estore: estore,
	}
}

type EnvRelationStore interface {
	AddUserToGroup(ctx context.Context, userid int64, envid int64) error
	RemoveUserFromGroup(ctx context.Context, userid int64, envid int64) error
	FindUserGroups(ctx context.Context, name string) ([]*enventity.UserEnviroment, error)
	FindUsersByGroup(ctx context.Context, name string) ([]*userentity.User, error)
}

type EnvsRel struct {
	erstore EnvRelationStore
}

func NewEnvsRel(erstore EnvRelationStore) *EnvsRel {
	return &EnvsRel{
		erstore: erstore,
	}
}
