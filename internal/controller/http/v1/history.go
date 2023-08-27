package v1

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
	"net/http"
)

type historyRoutes struct {
	historyService service.History
}

func newHistoryRoutes(g *echo.Group, historyService service.History) {
	r := historyRoutes{
		historyService: historyService,
	}
	g.POST("/report-link", r.reportLink) // POST as the file is created when called
}

type getHistoryInput struct {
	UserID string `json:"user_id" validate:"required,max=40"`
	Year   int    `json:"year" validate:"required"`
	Month  int    `json:"month" validate:"required"`
}

func (h *historyRoutes) reportLink(c echo.Context) error {
	var input getHistoryInput
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}
	filename, err := h.historyService.GetNotes(c.Request().Context(), service.GetHistoryInput{
		UserID: input.UserID,
		Year:   input.Year,
		Month:  input.Month,
	})
	if err != nil {
		if errors.Is(err, service.ErrUserNoData) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		return err
	}

	type response struct {
		UserID     string `json:"user_id"`
		ReportLink string `json:"report_link"`
	}

	return c.JSON(http.StatusOK, response{
		UserID:     input.UserID,
		ReportLink: fmt.Sprintf("http://%s/reports/%s", c.Request().Host, filename),
	})
}
