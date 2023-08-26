package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
	"net/http"
)

type userRoutes struct {
	userService service.User
}

func newUserRoutes(g *echo.Group, userService service.User) {
	r := &userRoutes{
		userService: userService,
	}
	g.POST("/segments", r.setSegments)
	g.GET("/", r.getSegments)
}

type setSegmentsUserInput struct {
	UserID      string   `json:"user_id" validate:"required,max=40"`
	SegmentsAdd []string `json:"segments_add" validate:"required"`
	SegmentsDel []string `json:"segments_del" validate:"required"`
	TTL         uint64   `json:"ttl" validate:"omitempty,min=1,max=18446744073709551615"`
}

func (u *userRoutes) setSegments(c echo.Context) error {
	var input setSegmentsUserInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	err := u.userService.SetSegments(c.Request().Context(), service.SetSegmentsUserInput{
		UserID:      input.UserID,
		SegmentsAdd: input.SegmentsAdd,
		SegmentsDel: input.SegmentsDel,
		TTL:         input.TTL,
	})
	if err != nil {
		if errors.Is(err, service.ErrSegmentNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return err
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}
	return c.NoContent(200)
}

func (u *userRoutes) getSegments(c echo.Context) error {
	return nil
}
