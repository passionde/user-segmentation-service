package pgdb

import "github.com/passionde/user-segmentation-service/pkg/postgres"

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// todo: реализовать интерфейс из repo.go
