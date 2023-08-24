package pgdb

import "github.com/passionde/user-segmentation-service/pkg/postgres"

type HistoryRepo struct {
	*postgres.Postgres
}

func NewHistoryRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// todo: реализовать интерфейс из repo.go
