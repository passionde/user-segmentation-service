package service

import (
	"context"
	"errors"
	"github.com/passionde/user-segmentation-service/internal/entity"
	"github.com/passionde/user-segmentation-service/internal/repo"
	"github.com/passionde/user-segmentation-service/internal/repo/repoerrs"
)

type SegmentService struct {
	segmentRepo repo.Segment
	historyRepo repo.History
}

func NewSegmentService(segmentRepo repo.Segment, historyRepo repo.History) *SegmentService {
	return &SegmentService{
		segmentRepo: segmentRepo,
		historyRepo: historyRepo,
	}
}

func (s *SegmentService) CreateSegment(ctx context.Context, input CreateSegmentInput) error {
	err := s.segmentRepo.CreateSegment(ctx, input.Slug)
	// todo реализовать метод добавления процента пользователей
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return ErrSegmentAlreadyExists
		}
		return ErrCannotCreateSegment
	}
	return nil
}

func (s *SegmentService) DeleteSegment(ctx context.Context, input SegmentInput) error {
	usersID, err := s.segmentRepo.GetUsersInSegment(ctx, input.Slug)
	if err != nil {
		return err
	}

	err = s.segmentRepo.DeleteSegment(ctx, input.Slug)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrSegmentNotFound
		}
		return err
	}
	return s.historyRepo.AddNotes(ctx, cookNotesSegment(usersID, input.Slug))
}

func cookNotesSegment(usersID []string, segment string) []entity.History {
	notes := make([]entity.History, 0, len(usersID))
	for _, userID := range usersID {
		notes = append(notes, entity.History{
			UserID:      userID,
			SegmentSlug: segment,
			Type:        entity.OperationTypeSegmentDelete,
		})
	}
	return notes
}
