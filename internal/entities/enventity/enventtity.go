package enventity

import "time"

const (
	GROUP = iota
	ORGANIZATION
	PROJECT
	COMMUNITY
)

type UserEnviroment struct {
	ID        int64
	Name      string
	EnvType   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
}
