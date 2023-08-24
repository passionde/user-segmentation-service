package service

import "github.com/passionde/user-segmentation-service/internal/repo"

type UserService struct {
	userRepo repo.User
}

func NewUserService(userRepo repo.User) *UserService {
	return &UserService{userRepo: userRepo}
}

// todo: реализовать методы сервиса
