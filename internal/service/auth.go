package service

import "github.com/passionde/user-segmentation-service/internal/repo"

type AuthService struct {
	authRepo repo.History
}

func NewAuthService(authRepo repo.Auth) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (a *AuthService) ParseKey(string) (int, error) {
	// todo: реализовать проверку ключа API
	return 0, nil
}

// todo: реализовать методы сервиса
