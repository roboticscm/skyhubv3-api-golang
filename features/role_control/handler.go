package role_control

import (
	"backend/system/db"
	. "backend/system/error"
	"backend/system/gbfunc"
	"net/http"

	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	menuPath := c.QueryParam("menuPath")
	depId := c.QueryParam("depId")

	if menuPath == "" {
		return Errord400(c, "SYS.MSG.MISSING_MENU_PATH", "")
	}

	if depId == "" {
		return Errord400(c, "SYS.MSG.MISSING_DEPARTMENT_ID", "")
	}
	const sql = `SELECT * FROM find_roled_control(?, ?, ?) as json`
	json, err := db.SelectJson(sql, &depId, &menuPath, &userId)

	if err != nil {
		return LoadObjectError(c, err, "ROLE_CONTROL")
	}

	c.String(http.StatusOK, json.(string))
	return nil
}
