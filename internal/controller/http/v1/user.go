package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/passionde/user-segmentation-service/internal/service"
)

type userRoutes struct {
	userService service.User
}

func newUserRoutes(g *echo.Group, userService service.User) {
	//r := &userRoutes{
	//	userService: userService,
	//}
	//
	//g.POST("/create", r.create)
	//g.POST("/deposit", r.deposit) // POST, а не PUT, потому что неидемпотентно
	//g.POST("/withdraw", r.withdraw)
	//g.POST("/transfer", r.transfer)
	//g.GET("/", r.getBalance)
}
