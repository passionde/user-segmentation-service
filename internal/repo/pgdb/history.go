package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/passionde/user-segmentation-service/internal/entity"
	"github.com/passionde/user-segmentation-service/pkg/postgres"
)

type HistoryRepo struct {
	*postgres.Postgres
}

func NewHistoryRepo(pg *postgres.Postgres) *HistoryRepo {
	return &HistoryRepo{pg}
}

func (h *HistoryRepo) AddNotes(ctx context.Context, notes []entity.History) error {
	b := h.Builder.Insert("history").Columns("user_id", "segment_slug", "type")
	for _, note := range notes {
		b = b.Values(note.UserID, note.SegmentSlug, note.Type)
	}
	sql, args, _ := b.ToSql()

	err := h.Pool.QueryRow(ctx, sql, args...).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("HistoryRepo.AddNote - s.Pool.QueryRow: %v", err)
	}
	return nil
}

// todo: реализовать интерфейс из repo.go
