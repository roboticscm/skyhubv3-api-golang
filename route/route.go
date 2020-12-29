package route

import (
	authRoute "backend/system/features/authentication"
	localeResourceRoute "backend/system/features/locale_resource"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RootRoute struct {
}

func (route *RootRoute) RegisterRoute(server *echo.Echo) {
	server.Static("/css", "css")

	server.GET("/", indexHandler)
	// Register language rest api group
	langRoute := localeResourceRoute.Route{Path: "locale-resource"}
	langRoute.RegisterRoute(server)

	// Register authentication rest api group
	authRoute := authRoute.Route{Path: "auth"}
	authRoute.RegisterRoute(server)

	server.Pre(middleware.AddTrailingSlash())
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Index Route Works!")
}
