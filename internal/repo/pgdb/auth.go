package pgdb

import "github.com/passionde/user-segmentation-service/pkg/postgres"

type AuthRepo struct {
	*postgres.Postgres
}

func NewAuthRepo(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg}
}

// todo: реализовать интерфейс из repo.go
