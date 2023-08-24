package service

import (
	"github.com/passionde/user-segmentation-service/internal/repo"
)

type CreateSegmentInput struct {
	// todo: Описать структуру данных для создания сегмента
}

type DeleteSegmentInput struct {
	// todo: Описать структуру данных для удаления сегмента
}

type Segment interface {
	// todo: Описать методы сегментов
}

type SetSegmentsUserInput struct {
	// todo: Описать структуру данных для установки сегментов пользователю
}

type GetSegmentsUserInput struct {
	// todo: Описать структуру данных для получения сегментов пользователя
}

type GetSegmentsUserOutput struct {
	// todo: Описать структуру данных для получения сегментов пользователя
}

type User interface {
	// todo: Описать методы пользователя
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
	ParseKey(string) (int, error)
}

type Services struct {
	User    User
	Segment Segment
	History History
	Auth    Auth
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		User:    NewUserService(deps.Repos.User),
		Segment: NewSegmentService(deps.Repos.Segment),
		History: NewHistoryService(deps.Repos.History),
		Auth:    NewAuthService(deps.Repos.Auth),
	}
}
