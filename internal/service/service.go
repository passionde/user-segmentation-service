package service

import (
	"context"
	"github.com/passionde/user-segmentation-service/internal/entity"
	"github.com/passionde/user-segmentation-service/internal/repo"
	"github.com/passionde/user-segmentation-service/pkg/csvwriter"
	"github.com/passionde/user-segmentation-service/pkg/secure"
)

type CreateSegmentInput struct {
	Slug            string
	PercentageUsers int
}

type SegmentInput struct {
	Slug string
}

type Segment interface {
	CreateSegment(ctx context.Context, input CreateSegmentInput) error
	DeleteSegment(ctx context.Context, input SegmentInput) error
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
	UserID string
	Month  int
	Year   int
}

type History interface {
	AddNotes(ctx context.Context, notes []entity.History) error
	GetNotes(ctx context.Context, input GetHistoryInput) (string, error)
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
	CSVWrite  csvwriter.CSVWriter
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		User:    NewUserService(deps.Repos.User, deps.Repos.History),
		Segment: NewSegmentService(deps.Repos.Segment, deps.Repos.History, deps.Repos.User),
		History: NewHistoryService(deps.Repos.History, deps.CSVWrite),
		Auth:    NewAuthService(deps.Repos.Auth, deps.APISecure),
	}
}
