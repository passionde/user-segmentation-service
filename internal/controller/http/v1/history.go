package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
)

type historyRoutes struct {
	historyService service.History
}

func newHistoryRoutes(g *echo.Group, historyService service.History) {
}
