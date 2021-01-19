package gbfunc

import (
	"backend/system/keys"
	"backend/system/lib"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
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
		d, _ := lib.ToInt64(maps[0]["date"])
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
			userId, err := lib.ToInt64(claims["UserId"])

			if err != nil {
				return -1, err
			}

			return userId, nil
		}
	}

	return -1, nil
}

func GetUserId(c echo.Context) int64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userIdStr := claims["userId"].(string)

	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	return userId
}

func ToJson(source string) interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(source), &result)

	if result != nil {
		return result
	}

	var result2 []map[string]interface{}
	json.Unmarshal([]byte(source), &result2)
	return result2
}

func MakeUpdate(c echo.Context, model interface{}) {
	var userId int64 = GetUserId(c)
	var now int64 = GetCurrentMillis()

	reflect.ValueOf(model).Elem().FieldByName("UpdatedBy").Set(reflect.ValueOf(&userId))
	reflect.ValueOf(model).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(&now))
}

func MakeSave(c echo.Context, model interface{}) {
	var userId int64 = GetUserId(c)
	var now int64 = GetCurrentMillis()
	var disabled = false
	reflect.ValueOf(model).Elem().FieldByName("CreatedBy").Set(reflect.ValueOf(&userId))
	reflect.ValueOf(model).Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(&now))
	reflect.ValueOf(model).Elem().FieldByName("Disabled").Set(reflect.ValueOf(&disabled))
}

func MakeDelete(c echo.Context, model interface{}) {
	var userId int64 = GetUserId(c)
	var now int64 = GetCurrentMillis()

	reflect.ValueOf(model).Elem().FieldByName("DeletedBy").Set(reflect.ValueOf(&userId))
	reflect.ValueOf(model).Elem().FieldByName("DeletedAt").Set(reflect.ValueOf(&now))
}

func IsTextValueDuplicated(tableName string, columnName string, value string, id int64) (bool, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	const sql = `SELECT * FROM is_text_value_duplicated(?, ?, ?, ?) as "isDuplicated"`
	if _, err := o.Raw(sql, tableName, columnName, value, id).Values(&maps); err != nil {
		return false, errors.New("SYS.MSG.LOAD_OBJECT_ERROR")
	}
	result, err := strconv.ParseBool(maps[0]["isDuplicated"].(string))

	return result, err
}

func IsTextValueExisted(tableName string, columnName string, value string) (bool, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	const sql = `SELECT * FROM is_text_value_existed(?, ?, ?) as "isExisted"`
	if _, err := o.Raw(sql, tableName, columnName, value).Values(&maps); err != nil {
		return false, errors.New("SYS.MSG.LOAD_OBJECT_ERROR")
	}
	result, err := strconv.ParseBool(maps[0]["isExisted"].(string))

	return result, err
}
