package menu

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	depId := c.QueryParam("depId")

	if depId == "" {
		return Errord400(c, "SYS.MSG.MISSING_DEPARTMENT_ID", "")
	}

	o := orm.NewOrm()
	menu := []Menu{}

	const sql = `SELECT * FROM find_menu(?, ?)`
	if _, err := o.Raw(sql, userId, depId).QueryRows(&menu); err != nil {
		return LoadObjectError(c, err, "LANGUAGE")
	}

	c.JSON(http.StatusOK, menu)
	return nil
}
