package service

import (
	"context"
	"errors"
	"github.com/passionde/user-segmentation-service/internal/repo"
	"github.com/passionde/user-segmentation-service/internal/repo/repoerrs"
)

type UserService struct {
	userRepo repo.User
}

func NewUserService(userRepo repo.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) SetSegments(ctx context.Context, input SetSegmentsUserInput) error {
	err := u.userRepo.SetSegments(
		ctx,
		input.UserID,
		input.SegmentsAdd,
		input.SegmentsDel,
	)
	if err != nil {
		if errors.Is(err, repoerrs.ErrSegmentsNotExist) {
			return ErrSegmentNotFound
		}
	}
	return err
	// todo реализовать TTL
}

func (u *UserService) GetSegments(ctx context.Context, input GetSegmentsUserInput) ([]string, error) {
	segments, err := u.userRepo.GetSegments(ctx, input.UserID)
	if err != nil {
		if errors.Is(err, repoerrs.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return segments, nil
}
