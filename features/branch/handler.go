package branch

import (
	"backend/system/db"
	. "backend/system/error"
	"backend/system/gbfunc"
	"net/http"

	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	fromOrgType := c.QueryParam("fromOrgType")
	toOrgType := c.QueryParam("toOrgType")
	includeDisabled := c.QueryParam("includeDisabled")
	includeDeleted := c.QueryParam("includeDeleted")

	const sql = `SELECT * FROM find_branch_tree(?, ?, ?, ?, ?) as "json"`
	json, err := db.SelectJson(sql, &userId, &fromOrgType, &toOrgType, &includeDeleted, &includeDisabled)

	if err != nil {
		return LoadObjectError(c, err, "BRANCH")
	}

	c.String(http.StatusOK, json.(string))
	return nil
}
