package table_util

import (
	"backend/system/db"
	. "backend/system/error"
	"backend/system/gbfunc"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func findSimpleListHandler(c echo.Context) error {
	tableName := c.QueryParam("tableName")
	columns := c.QueryParam("columns")
	orderBy := c.QueryParam("orderBy")
	page := c.QueryParam("page")
	pageSize := c.QueryParam("pageSize")
	onlyMe := c.QueryParam("onlyMe")
	includeDisabled := c.QueryParam("includeDisabled")
	userId := gbfunc.GetUserId(c)

	if tableName == "" {
		return Errord400(c, "SYS.MSG.MISSING_TABLE_NAME", "")
	}

	if columns == "" {
		return Errord400(c, "SYS.MSG.MISSING_COLUMN_LIST", "")
	}

	if orderBy == "" {
		return Errord400(c, "SYS.MSG.MISSING_ORDER_BY_COLUMN", "")
	}

	if page == "" {
		return Errord400(c, "SYS.MSG.MISSING_PAGE", "")
	}

	if pageSize == "" {
		return Errord400(c, "SYS.MSG.MISSING_PAGE_SIZE", "")
	}

	if onlyMe == "" {
		return Errord400(c, "SYS.MSG.MISSING_ONLY_ME", "")
	}

	if includeDisabled == "" {
		return Errord400(c, "SYS.MSG.MISSING_INCLUDE_DISABLED", "")
	}

	const sql = `SELECT * FROM find_simple_list(?, ?, ?, ?, ?, ?, ?, ?) as json`

	json, err := db.SelectJson(sql, tableName, columns, orderBy, page, pageSize, onlyMe, userId, includeDisabled)

	if err != nil {
		return LoadObjectError(c, err, "TABLE_UTIL")
	}

	// c.JSON(http.StatusOK, json)
	c.String(http.StatusOK, json.(string))
	return nil
}

func getOneHandler(c echo.Context) error {
	tableName := c.QueryParam("tableName")
	id := c.Param("id")
	if tableName == "" {
		return Errord400(c, "SYS.MSG.MISSING_TABLE_NAME", "")
	}

	if id == "" {
		return Errord400(c, "SYS.MSG.MISSING_ID", "")
	}

	const sql = `SELECT * FROM get_one_by_id(?, ?) as json`

	json, err := db.SelectJson(sql, tableName, id)

	if err != nil {
		return LoadObjectError(c, err, "TABLE_UTIL")
	}

	// c.JSON(http.StatusOK, json)
	c.String(http.StatusOK, json.(string))
	return nil
}

func hasAnyDeletedRecordHandler(c echo.Context) error {
	tableName := c.QueryParam("tableName")
	onlyMe := c.QueryParam("onlyMe")
	userId := gbfunc.GetUserId(c)

	if tableName == "" {
		return Errord400(c, "SYS.MSG.MISSING_TABLE_NAME", "")
	}

	if onlyMe == "" {
		onlyMe = "false"
	}

	const sql = `SELECT * FROM has_any_deleted_record(?, ?, ?) as json`
	json, err := db.SelectJson(sql, tableName, onlyMe, userId)

	if err != nil {
		return LoadObjectError(c, err, "TABLE_UTIL")
	}

	// c.JSON(http.StatusOK, json)
	c.String(http.StatusOK, json.(string))
	return nil
}

func restoreOrForeverDeleteHandler(c echo.Context) error {
	tableName := c.QueryParam("tableName")
	companyId := c.QueryParam("companyId")
	branchId := c.QueryParam("branchId")
	menuPath := c.QueryParam("menuPath")
	ipClient := c.QueryParam("ipClient")
	device := c.QueryParam("device")
	os := c.QueryParam("os")
	browser := c.QueryParam("browser")
	fieldName := c.QueryParam("fieldName")
	deleteIds := c.QueryParam("deleteIds")
	restoreIds := c.QueryParam("restoreIds")
	reason := c.QueryParam("reason")

	userId := gbfunc.GetUserId(c)

	if tableName == "" {
		return Errord400(c, "SYS.MSG.MISSING_TABLE_NAME", "")
	}

	deleteIdsRef := &deleteIds
	if deleteIds == "" {
		deleteIdsRef = nil
	}

	restoreIdsRef := &restoreIds
	if restoreIds == "" {
		restoreIdsRef = nil
	}

	reasonRef := &reason
	if reason == "" {
		reasonRef = nil
	}

	if companyId != "" {
		const sql = `SELECT * FROM restore_or_forever_delete(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) as json`
		json, err := db.SelectJson(sql, tableName, &deleteIdsRef, &restoreIdsRef, userId, companyId, branchId, menuPath, ipClient, device, os, browser, &reasonRef, fieldName)

		if err != nil {
			return LoadObjectError(c, err, "TABLE_UTIL")
		}
		c.JSON(http.StatusOK, json)
	} else {
		const sql = `SELECT * FROM restore_or_forever_delete(?, ?, ?, ?) as json`
		json, err := db.SelectJson(sql, tableName, &deleteIdsRef, &restoreIdsRef, userId)

		if err != nil {
			return LoadObjectError(c, err, "TABLE_UTIL")
		}
		c.String(http.StatusOK, json.(string))
		// c.JSON(http.StatusOK, json)
	}
	return nil
}

func findDeletedRecordsHandler(c echo.Context) error {
	tableName := c.QueryParam("tableName")
	columns := c.QueryParam("columns")
	onlyMe := c.QueryParam("onlyMe")
	userId := gbfunc.GetUserId(c)

	if tableName == "" {
		return Errord400(c, "SYS.MSG.MISSING_TABLE_NAME", "")
	}

	if columns == "" {
		return Errord400(c, "SYS.MSG.MISSING_COLUMN_LIST", "")
	}

	if onlyMe == "" {
		onlyMe = "false"
	}

	const sql = `SELECT * FROM find_deleted_records(?, ?, ?, ?) as json`
	json, err := db.SelectJson(sql, tableName, columns, onlyMe, userId)

	if err != nil {
		return LoadObjectError(c, err, "TABLE_UTIL")
	}
	c.String(http.StatusOK, json.(string))
	return nil
}

func softDeleteManyHandler(c echo.Context) error {
	tableName := c.QueryParam("tableName")
	ids := c.QueryParam("ids")
	userId := gbfunc.GetUserId(c)

	if tableName == "" {
		return Errord400(c, "SYS.MSG.MISSING_TABLE_NAME", "")
	}

	if ids == "" {
		return Errord400(c, "SYS.MSG.MISSING_ID_LIST", "")
	}

	var sql = "UPDATE " + tableName + " "
	sql += `	
		SET deleted_by = ?, deleted_at = ?
		WHERE id IN
	`
	sql += " (" + ids + ")"
	o := orm.NewOrm()
	now := gbfunc.GetCurrentMillis()
	res, err := o.Raw(sql, userId, now).Exec()

	if err != nil {
		return UpdateObjectError(c, err, "TABLE_UTIL")
	}
	num, _ := res.RowsAffected()
	c.JSON(http.StatusOK, map[string]interface{}{
		"deletedRows": num,
	})

	return nil
}
