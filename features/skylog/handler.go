package skylog

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"backend/system/models"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	menuPath := c.QueryParam("menuPath")
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")

	var sql = `
        SELECT l.id, l.created_at as date, a.username as user, l.reason, l.description, l.short_description, '' as view
        FROM sky_log l
        INNER JOIN account a ON a.id = l.created_by
        WHERE l.menu_path = ?
    `
	param := []interface{}{menuPath}
	if startDate != "" {
		sql += " AND l.created_at >= ? "
		param = append(param, startDate)
	}

	if endDate != "" {
		sql += " AND l.created_at <= ? "
		param = append(param, endDate)
	}

	sql += " ORDER BY l.created_at DESC"

	skylogs := []SkyLog{}

	o := orm.NewOrm()
	if _, err := o.Raw(sql, param...).QueryRows(&skylogs); err != nil {
		return LoadObjectError(c, err, "SKYLOG")
	}

	c.JSON(http.StatusOK, skylogs)
	return nil
}

func saveHandler(c echo.Context) error {
	skylog := &models.SkyLog{}

	if err := c.Bind(skylog); err != nil {
		return BindObjectError(c, err, "SKYLOG")
	}

	o := orm.NewOrm()
	gbfunc.MakeSave(c, skylog)
	id, err := o.Insert(skylog)

	if err != nil {
		return SaveObjectError(c, err, "SKYLOG")
	}

	skylog.Id = id
	err = o.Read(skylog)
	if err != nil {
		return LoadObjectError(c, err, "SKYLOG")
	}
	c.JSON(http.StatusOK, skylog)
	return nil
}
