package pgdb

import "github.com/passionde/user-segmentation-service/pkg/postgres"

type SegmentRepo struct {
	*postgres.Postgres
}

func NewSegmentRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// todo: реализовать интерфейс из repo.go
