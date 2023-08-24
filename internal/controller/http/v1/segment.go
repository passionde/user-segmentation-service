package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
)

type segmentRoutes struct {
	segmentService service.Segment
}

func newSegmentRoutes(g *echo.Group, segmentService service.Segment) {
}
