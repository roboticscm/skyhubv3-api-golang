package route

import (
	"backend/system/db"
	authRoute "backend/system/features/authentication"
	branchRoute "backend/system/features/branch"
	departmentRoute "backend/system/features/department"
	languageRoute "backend/system/features/language"
	localeResourceRoute "backend/system/features/locale_resource"
	menuRoute "backend/system/features/menu"
	menuControlRoute "backend/system/features/menu_control"
	menuHistoryRoute "backend/system/features/menu_history"
	roleRoute "backend/system/features/role"
	roleControlRoute "backend/system/features/role_control"
	searchUtilRoute "backend/system/features/search_util"
	skylogRoute "backend/system/features/skylog"
	tableUtilRoute "backend/system/features/table_util"
	userSettingsRoute "backend/system/features/user_settings"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RootRoute struct {
}

func (route *RootRoute) RegisterRoute(server *echo.Echo, b *db.Broker) {
	server.Static("/css", "css")

	server.GET("/", indexHandler)
	server.GET("/notify/", b.NotifyHandler)

	// Register authentication rest api group
	authRoute.RegisterRoute("auth", server)

	// Register branch rest api group
	branchRoute.RegisterRoute("branch", server)

	// Register department rest api group
	departmentRoute.RegisterRoute("department", server)

	// Register language rest api group
	languageRoute.RegisterRoute("language", server)

	// Register language rest api group
	localeResourceRoute.RegisterRoute("locale-resource", server)

	// Register menu rest api group
	menuRoute.RegisterRoute("menu", server)

	// Register menu-control rest api group
	menuControlRoute.RegisterRoute("menu-control", server)

	// Register menu-history rest api group
	menuHistoryRoute.RegisterRoute("menu-history", server)

	// Register role rest api group
	roleRoute.RegisterRoute("role", server)

	// Register role-control rest api group
	roleControlRoute.RegisterRoute("role-control", server)

	// Register search-util rest api group
	searchUtilRoute.RegisterRoute("search-util", server)

	// Register skylog rest api group
	skylogRoute.RegisterRoute("skylog", server)

	// Register table-util rest api group
	tableUtilRoute.RegisterRoute("table-util", server)

	// Register user-settings rest api group
	userSettingsRoute.RegisterRoute("user-settings", server)

	server.Pre(middleware.AddTrailingSlash())
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Index Route Works!")
}
