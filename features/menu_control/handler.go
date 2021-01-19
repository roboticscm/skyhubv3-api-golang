package menu_control

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"backend/system/models"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	menuPath := c.QueryParam("menuPath")

	if menuPath == "" {
		return Errord400(c, "SYS.MSG.MISSING_MENU_PATH", "")
	}

	o := orm.NewOrm()
	menuControls := []MenuControl{}

	if _, err := o.Raw(`SELECT * FROM find_menu_control(?)`, menuPath).QueryRows(&menuControls); err != nil {
		return LoadObjectError(c, err, "MENU_CONTROL")
	}

	c.JSON(http.StatusOK, menuControls)
	return nil
}

func saveOrDeleteHandler(c echo.Context) error {
	menuControlBody := &MenuControlBody{}
	if err := c.Bind(menuControlBody); err != nil {
		return BindObjectError(c, err, "MENU_CONTROL")
	}

	o := orm.NewOrm()
	menu := []models.Menu{}

	if _, err := o.Raw(`SELECT id FROM menu WHERE path = ?`, menuControlBody.MenuPath).QueryRows(&menu); err != nil {
		return LoadObjectError(c, err, "MENU")
	}

	if len(menu) > 0 {
		for _, control := range menuControlBody.MenuControls {
			menuControls := []models.MenuControl{}
			const sql = `SELECT * FROM menu_control WHERE menu_id = ? AND control_id = ? `
			if _, err := o.Raw(sql, menu[0].Id, control.ControlId).QueryRows(&menuControls); err != nil {
				return LoadObjectError(c, err, "MENU_CONTROL")
			}

			if len(menuControls) > 0 {
				if !*control.Checked {
					if _, err := o.Delete(&menuControls[0]); err == nil {
						fmt.Println(err)
					}
				}
			} else {
				if *control.Checked {
					menuControl := models.MenuControl{MenuId: &menu[0].Id, ControlId: &control.ControlId}
					gbfunc.MakeSave(c, &menuControl)
					id, err := o.Insert(&menuControl)
					if err == nil {
						fmt.Println(id)
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK, map[string]string{"message": "success"})
	return nil
}
