package user_settings

import (
	"backend/system/features/authentication"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(path string, server *echo.Echo) {
	group := server.Group(path)
	group.GET("/initial/", getInitialHandler, authentication.IsAuthenticated())
	group.POST("/", saveUserSettingsHandler, authentication.IsAuthenticated())
	group.GET("/", getUserSettingsHandler, authentication.IsAuthenticated())
}
