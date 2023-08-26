package service

import (
	"context"
	"github.com/passionde/user-segmentation-service/internal/repo"
	"github.com/passionde/user-segmentation-service/pkg/secure"
)

type CreateSegmentInput struct {
	Slug            string
	PercentageUsers int
}

type DeleteSegmentInput struct {
	Slug string
}

type Segment interface {
	CreateSegment(ctx context.Context, input CreateSegmentInput) error
	DeleteSegment(ctx context.Context, input DeleteSegmentInput) error
}

type SetSegmentsUserInput struct {
	UserID      string
	SegmentsAdd []string
	SegmentsDel []string
	TTL         uint64
}

type GetSegmentsUserInput struct {
	UserID string
}

type User interface {
	SetSegments(ctx context.Context, input SetSegmentsUserInput) error
	GetSegments(ctx context.Context, input GetSegmentsUserInput) ([]string, error)
}

type GetHistoryInput struct {
	// todo: Описать структуру данных для получения истории пользователя
}

type GetHistoryOutput struct {
	// todo: Описать структуру данных для получения истории пользователя
}

type History interface {
	// todo: Описать методы истории
}

type Auth interface {
	TokenExist(ctx context.Context, token string) (int, error)
	GenerateToken(ctx context.Context) (int, string, error)
}

type Services struct {
	User    User
	Segment Segment
	History History
	Auth    Auth
}

type ServicesDependencies struct {
	Repos     *repo.Repositories
	APISecure secure.APISecure
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		User:    NewUserService(deps.Repos.User),
		Segment: NewSegmentService(deps.Repos.Segment),
		History: NewHistoryService(deps.Repos.History),
		Auth:    NewAuthService(deps.Repos.Auth, deps.APISecure),
	}
}
