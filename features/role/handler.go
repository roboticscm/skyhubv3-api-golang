package role

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"backend/system/lib"
	"backend/system/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func saveOrUpdateHandler(c echo.Context) error {
	roleBody := &RoleBody{}

	if err := c.Bind(roleBody); err != nil {
		return BindObjectError(c, err, "ROLE")
	}

	if roleBody.Code != nil {
		*roleBody.Code = strings.TrimSpace(*roleBody.Code)
	}

	if roleBody.Name != nil {
		*roleBody.Name = strings.TrimSpace(*roleBody.Name)
	}

	o := orm.NewOrm()
	if roleBody.Id != nil {
		isDuplicated, _ := gbfunc.IsTextValueDuplicated("role", "code", *roleBody.Code, *roleBody.Id)
		if isDuplicated {
			return DuplicatedError(c, "code")
		}

		isDuplicated, _ = gbfunc.IsTextValueDuplicated("role", "name", *roleBody.Name, *roleBody.Id)
		if isDuplicated {
			return DuplicatedError(c, "name")
		}

		role := models.Role{Id: *roleBody.Id}
		err := o.Read(&role)
		if err != nil {
			return LoadObjectError(c, err, "ROLE")
		}
		role.Code = roleBody.Code
		role.Name = roleBody.Name
		role.Sort = roleBody.Sort
		*role.OrgId, _ = lib.ToInt64(roleBody.OrgId)
		gbfunc.MakeUpdate(c, &role)
		if num, err := o.Update(&role); err == nil {
			c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Updated %v", num)})
		}
	} else {
		isExisted, _ := gbfunc.IsTextValueExisted("role", "code", *roleBody.Code)
		if isExisted {
			return ExistedError(c, "code")
		}

		isExisted, _ = gbfunc.IsTextValueExisted("role", "name", *roleBody.Name)
		if isExisted {
			return ExistedError(c, "name")
		}
		orgId, _ := lib.ToInt64(roleBody.OrgId)
		role := models.Role{Code: roleBody.Code, Name: roleBody.Name, Sort: roleBody.Sort, OrgId: &orgId}
		gbfunc.MakeSave(c, &role)
		if inserted, err := o.Insert(&role); err == nil {
			c.JSON(http.StatusOK, inserted)
		}
	}

	return nil
}
