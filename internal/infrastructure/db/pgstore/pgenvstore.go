package pgstore

import (
	"backend2/internal/entities/enventity"
	"backend2/internal/usecases/app/repos/envrepo"
	"context"
	"database/sql"
	"time"
)

var _ envrepo.EnvStore = &Envs{}

type Env struct {
	ID        int64
	Name      string
	EnvType   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
}

type Envs struct {
	db *sql.DB
}

func NewEnvs(db *sql.DB) *Envs {
	en := &Envs{
		db: db,
	}
	return en
}

func (e *Envs) CreateEnv(ctx context.Context, env *envrepo.Enviroment) (*enventity.UserEnviroment, error) {
	en := &enventity.UserEnviroment{}
	return en, nil
}

func (e *Envs) FindEnv(ctx context.Context, id int64) (*enventity.UserEnviroment, error) {
	en := &enventity.UserEnviroment{}
	return en, nil
}

func (e *Envs) DeleteEnv(ctx context.Context, id int64) error {
	return nil
}
