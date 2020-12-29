package authentication

import (
	mdw "backend/system/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Route struct {
	Path string
}

func (route *Route) RegisterRoute(server *echo.Echo) {
	group := server.Group(route.Path)
	group.POST("/login/", loginHandler, middleware.BasicAuth(mdw.BasicAuth))
	group.POST("/refresh-token/", refreshTokenHandler)
	group.DELETE("/logout/", logoutHandler)

}
