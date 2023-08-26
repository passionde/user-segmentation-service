package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/passionde/user-segmentation-service/internal/repo/repoerrs"
	"github.com/passionde/user-segmentation-service/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (u *UserRepo) GetSegments(ctx context.Context, userID string) ([]string, error) {

}

func (u *UserRepo) SetSegments(ctx context.Context, userID string, segmentsAdd, segmentsDel []string) error {
	userID, err := u.createUserIfNotExist(ctx, userID)
	if err != nil {
		return fmt.Errorf("UserRepo.SetSegments - u.createUserIfNotExist: %v", err)
	}

	ok, err := u.checkExistSegmentsSlug(ctx, segmentsAdd)
	if err != nil {
		return fmt.Errorf("UserRepo.SetSegments - u.checkExistSegmentsSlug: %v", err)
	}
	if !ok {
		return repoerrs.ErrSegmentsNotExist
	}

	if err := u.addSegmentsUser(ctx, userID, segmentsAdd); err != nil {
		return fmt.Errorf("UserRepo.SetSegments - u.addSegmentsUser: %v", err)
	}
	if err := u.delSegmentsUser(ctx, userID, segmentsDel); err != nil {
		return fmt.Errorf("UserRepo.SetSegments - u.addSegmentsUser: %v", err)
	}
	return nil
}

func (u *UserRepo) createUserIfNotExist(ctx context.Context, userID string) (string, error) {
	sql, args, _ := u.Builder.
		Insert("users").
		Columns("user_id").
		Values(userID).
		Suffix("ON CONFLICT (user_id) DO NOTHING RETURNING user_id").
		ToSql()
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return userID, nil
		}
		return "", fmt.Errorf("UserRepo.CreateUserIfNotExist - u.Pool.QueryRow: %v", err)
	}
	return userID, nil
}

func (u *UserRepo) checkExistSegmentsSlug(ctx context.Context, segmentSlugs []string) (bool, error) {
	sql, args, _ := u.Builder.
		Select("COUNT(*) AS count_found_segments").
		From("segments").
		Where(squirrel.Eq{"slug": segmentSlugs}).
		ToSql()

	var count int
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("UserRepo.CheckExistSegmentsSlug - u.Pool.QueryRow: %v", err)
	}

	return count == len(segmentSlugs), nil
}

func (u *UserRepo) addSegmentsUser(ctx context.Context, userID string, segmentSlugs []string) error {
	b := u.Builder.Insert("user_segments").Columns("user_id", "segment_slug")
	for _, segment := range segmentSlugs {
		b = b.Values(userID, segment)
	}
	sql, args, _ := b.Suffix("ON CONFLICT (user_id, segment_slug) DO NOTHING").ToSql()
	err := u.Pool.QueryRow(ctx, sql, args...).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("UserRepo.addSegmentsUser - u.Pool.QueryRow: %v", err)
	}
	return nil
}

func (u *UserRepo) delSegmentsUser(ctx context.Context, userID string, segmentSlugs []string) error {
	sql, args, _ := u.Builder.Delete("user_segments").Where(squirrel.And{
		squirrel.Eq{"user_id": userID},
		squirrel.Eq{"segment_slug": segmentSlugs},
	}).ToSql()

	err := u.Pool.QueryRow(ctx, sql, args...).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("UserRepo.addSegmentsUser - u.Pool.QueryRow: %v", err)
	}
	return nil
}
