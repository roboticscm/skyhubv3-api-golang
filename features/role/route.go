package role

import (
	"backend/system/features/authentication"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(path string, server *echo.Echo) {
	group := server.Group(path)
	group.POST("/", saveOrUpdateHandler, authentication.IsAuthenticated())
}
