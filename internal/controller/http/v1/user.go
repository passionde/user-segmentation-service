package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
)

type userRoutes struct {
	userService service.User
}

func newUserRoutes(g *echo.Group, userService service.User) {
	r := &userRoutes{
		userService: userService,
	}
	g.POST("/set-segments", r.setSegments)
	g.GET("/", r.getSegments)
}

func (u *userRoutes) setSegments(c echo.Context) error {
	return nil
}

func (u *userRoutes) getSegments(c echo.Context) error {
	return nil
}
