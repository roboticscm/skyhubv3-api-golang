package department

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	branchId := c.QueryParam("branchId")

	if branchId == "" {
		return Errord400(c, "SYS.MSG.MISSING_BRANCH_ID", "")
	}

	o := orm.NewOrm()
	departments := []Department{}

	sql := "SELECT * FROM find_department(?, ?)"
	if _, err := o.Raw(sql, branchId, userId).QueryRows(&departments); err != nil {
		return LoadObjectError(c, err, "DEPARTMENT")
	}

	c.JSON(http.StatusOK, departments)
	return nil
}

func getLastHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	branchId := c.QueryParam("branchId")

	if branchId == "" {
		return Errord400(c, "SYS.MSG.MISSING_BRANCH_ID", "")
	}

	o := orm.NewOrm()

	departmentIds := []DepartmentId{}

	const sql = `SELECT * FROM get_last_department_id(?, ?) as dep_id`
	if _, err := o.Raw(sql, branchId, userId).QueryRows(&departmentIds); err != nil {
		return LoadObjectError(c, err, "DEPARTMENT")
	}

	c.JSON(http.StatusOK, departmentIds)
	return nil
}
