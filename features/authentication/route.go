package authentication

import (
	mdw "backend/system/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoute(path string, server *echo.Echo) {
	group := server.Group(path)
	group.POST("/login/", loginHandler, middleware.BasicAuth(mdw.BasicAuth))
	group.PUT("/change-pw/", changePasswordHandler, IsAuthenticated())
	group.POST("/refresh-token/", refreshTokenHandler)
	group.DELETE("/logout/", logoutHandler)

}
