package repo

import (
	"github.com/passionde/user-segmentation-service/internal/repo/pgdb"
	"github.com/passionde/user-segmentation-service/pkg/postgres"
)

type User interface {
	// todo: перенести методы
}

type Segment interface {
	// todo: перенести методы
}

type History interface {
	// todo: перенести методы
}

type Auth interface {
	// todo: перенести методы
}

type Repositories struct {
	User
	Segment
	History
	Auth
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User:    pgdb.NewUserRepo(pg),
		Segment: pgdb.NewSegmentRepo(pg),
		History: pgdb.NewHistoryRepo(pg),
		Auth:    pgdb.NewAuthRepo(pg),
	}
}
