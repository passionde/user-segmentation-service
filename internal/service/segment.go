package service

import (
	"context"
	"errors"
	"github.com/passionde/user-segmentation-service/internal/repo"
	"github.com/passionde/user-segmentation-service/internal/repo/repoerrs"
)

type SegmentService struct {
	segmentRepo repo.Segment
}

func NewSegmentService(segmentRepo repo.Segment) *SegmentService {
	return &SegmentService{segmentRepo: segmentRepo}
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

func (s *SegmentService) DeleteSegment(ctx context.Context, input DeleteSegmentInput) error {
	err := s.segmentRepo.DeleteSegment(ctx, input.Slug)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrSegmentNotFound
		}
	}
	return err
}
