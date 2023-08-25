package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
)

type segmentRoutes struct {
	segmentService service.Segment
}

func newSegmentRoutes(g *echo.Group, segmentService service.Segment) {
	r := segmentRoutes{
		segmentService: segmentService,
	}
	g.POST("/create", r.create)
	g.DELETE("/delete", r.delete)
}

func (s *segmentRoutes) create(c echo.Context) error {
	return nil
}

func (s *segmentRoutes) delete(c echo.Context) error {
	return nil
}
