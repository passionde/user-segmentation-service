package service

import "github.com/passionde/user-segmentation-service/internal/repo"

type HistoryService struct {
	historyRepo repo.History
}

func NewHistoryService(historyRepo repo.History) *HistoryService {
	return &HistoryService{historyRepo: historyRepo}
}

// todo: реализовать методы сервиса
