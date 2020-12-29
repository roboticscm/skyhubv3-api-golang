package gbfunc

import (
	"backend/system/convertor"
	"backend/system/keys"
	"errors"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetCurrentMillis() int64 {
	o := orm.NewOrm()
	sql := `SELECT date_generator() as date`

	var maps []orm.Params
	num, err := o.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		d, _ := convertor.StringToInt64(maps[0]["date"])
		return d
	}

	return -1
}

func GetCurrentUserId(c echo.Context) (int64, error) {
	auth := c.Request().Header.Get("Authorization")

	if auth != "" {
		token := strings.Replace(auth, "Bearer ", "", -1)
		if token != "" {
			j, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				return keys.VerifyKey, nil
			})

			if err != nil {
				return -1, errors.New("401. Unauthorized")
			}

			claims := j.Claims.(jwt.MapClaims)
			userId, err := convertor.StringToInt64(claims["UserId"])

			if err != nil {
				return -1, err
			}

			return userId, nil
		}
	}

	return -1, nil
}
