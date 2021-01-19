package menu_history

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"backend/system/models"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func saveHandler(c echo.Context) error {
	menuHistoryBody := &MenuHistoryBody{}
	if err := c.Bind(menuHistoryBody); err != nil {
		return BindObjectError(c, err, "MENU_HISTORY")
	}

	userId := gbfunc.GetUserId(c)

	o := orm.NewOrm()
	menuHistories := []models.MenuHistory{}

	const sql = `SELECT * FROM menu_history WHERE menu_id = ? AND dep_id = ? AND account_id = ?`
	if _, err := o.Raw(sql, menuHistoryBody.MenuId, menuHistoryBody.DepId, userId).QueryRows(&menuHistories); err != nil {
		return LoadObjectError(c, err, "MENU_HISTORY")
	}

	now := gbfunc.GetCurrentMillis()
	if len(menuHistories) > 0 {
		menuHistories[0].LastAccess = &now
		if _, err := o.Update(&menuHistories[0]); err != nil {
			return UpdateObjectError(c, err, "MENU_HISTORY")
		}
	} else {
		menuHistory := models.MenuHistory{MenuId: &menuHistoryBody.MenuId, DepId: &menuHistoryBody.DepId, AccountId: &userId, LastAccess: &now}
		gbfunc.MakeSave(c, &menuHistory)
		_, err := o.Insert(&menuHistory)
		if err != nil {
			return SaveObjectError(c, err, "MENU_HISTORY")
		}

	}
	c.JSON(http.StatusOK, map[string]string{"message": "success"})
	return nil
}
