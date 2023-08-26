package repo

import (
	"context"
	"github.com/passionde/user-segmentation-service/internal/repo/pgdb"
	"github.com/passionde/user-segmentation-service/pkg/postgres"
)

type User interface {
	SetSegments(ctx context.Context, userID string, segmentsAdd, segmentsDel []string) error
	GetSegments(ctx context.Context, userID string) ([]string, error)
}

type Segment interface {
	CreateSegment(ctx context.Context, slug string) error
	DeleteSegment(ctx context.Context, slug string) error
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
