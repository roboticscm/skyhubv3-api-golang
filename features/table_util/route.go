package table_util

import (
	"backend/system/features/authentication"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(path string, server *echo.Echo) {
	group := server.Group(path)
	group.GET("/simple-list/", findSimpleListHandler, authentication.IsAuthenticated())
	group.GET("/has-any-deleted-record/", hasAnyDeletedRecordHandler, authentication.IsAuthenticated())
	group.GET("/find-deleted-records/", findDeletedRecordsHandler, authentication.IsAuthenticated())
	group.PUT("/restore-or-forever-delete/", restoreOrForeverDeleteHandler, authentication.IsAuthenticated())
	group.GET("/:id/", getOneHandler, authentication.IsAuthenticated())
	group.DELETE("/", softDeleteManyHandler, authentication.IsAuthenticated())
}
