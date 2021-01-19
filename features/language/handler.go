package language

import (
	. "backend/system/error"
	. "backend/system/models"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	o := orm.NewOrm()
	languages := []Language{}

	sql := `SELECT *
		FROM language
		WHERE disabled = FALSE
			AND deleted_by IS NULL
		ORDER BY sort, name
	`
	if _, err := o.Raw(sql).QueryRows(&languages); err != nil {
		return LoadObjectError(c, err, "LANGUAGE")
	}

	c.JSON(http.StatusOK, languages)
	return nil
}
