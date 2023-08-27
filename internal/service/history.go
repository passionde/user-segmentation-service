package service

import (
	"context"
	"github.com/passionde/user-segmentation-service/internal/entity"
	"github.com/passionde/user-segmentation-service/internal/repo"
)

type HistoryService struct {
	historyRepo repo.History
}

func NewHistoryService(historyRepo repo.History) *HistoryService {
	return &HistoryService{historyRepo: historyRepo}
}

func (h *HistoryService) AddNotes(ctx context.Context, notes []entity.History) error {
	return h.historyRepo.AddNotes(ctx, notes)
}

// todo: реализовать методы сервиса
