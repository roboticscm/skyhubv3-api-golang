package locale_resource

import (
	"backend/system/features/authentication"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(path string, server *echo.Echo) {
	group := server.Group(path)
	group.GET("/get-initial/", getInitialHandler)
	group.GET("/", getHandler, authentication.IsAuthenticated())
	group.POST("/", saveHandler, authentication.IsAuthenticated())
	group.PUT("/", updateHandler, authentication.IsAuthenticated())
	group.DELETE("/", deleteHandler, authentication.IsAuthenticated())
	group.GET("/report/", generateReportHanlder, authentication.IsAuthenticated())
	group.POST("/report/", generatePdfHanlder, authentication.IsAuthenticated())
}
