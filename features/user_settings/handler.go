package user_settings

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"backend/system/lib"
	"strings"

	"backend/system/models"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func getInitialHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)

	const sql = `SELECT * FROM get_last_user_settings(?)`

	userSettings := []UserSetting{}
	o := orm.NewOrm()
	_, err := o.Raw(sql, userId).QueryRows(&userSettings)

	if err != nil {
		return LoadObjectError(c, err, "USER_SETTING")
	}

	c.JSON(http.StatusOK, userSettings)
	return nil
}

func saveUserSettingsHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	userSettingsBody := &UserSettingBody{}

	if err := c.Bind(userSettingsBody); err != nil {
		return BindObjectError(c, err, "USER_SETTING")
	}

	o := orm.NewOrm()
	for index, key := range userSettingsBody.Keys {
		qs := o.QueryTable("user_setting").Filter("key", key)
		if userSettingsBody.BranchId != nil {
			qs = qs.Filter("branchId", userSettingsBody.BranchId)
		} else {
			qs = qs.Filter("branchId__isnull", true)
		}

		if userSettingsBody.MenuPath != nil {
			qs = qs.Filter("menuPath", userSettingsBody.MenuPath)
		} else {
			qs = qs.Filter("menuPath__isnull", true)
		}

		if userSettingsBody.ElementId != nil {
			qs = qs.Filter("elementId", userSettingsBody.ElementId)
		} else {
			qs = qs.Filter("elementId__isnull", true)
		}

		updatedRows, err := qs.Update(orm.Params{
			"value": userSettingsBody.Values[index],
		})

		if err != nil {
			return UpdateObjectError(c, err, "USER_SETTING")
		}

		if updatedRows == 0 {
			branchId, _ := lib.ToInt64(userSettingsBody.BranchId)
			userSetting := models.UserSetting{
				BranchId:  &branchId,
				AccountId: &userId,
				MenuPath:  userSettingsBody.MenuPath,
				ElementId: userSettingsBody.ElementId,
				Key:       &key,
				Value:     &userSettingsBody.Values[index],
			}

			_, err := o.Insert(&userSetting)

			if err != nil {
				return SaveObjectError(c, err, "USER_SETTING")
			}
		}
	}

	c.JSON(http.StatusOK, map[string]string{"message": "success"})
	return nil
}

func getUserSettingsHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	branchId := c.QueryParam("branchId")
	menuPath := c.QueryParam("menuPath")
	elementId := c.QueryParam("elementId")
	key := c.QueryParam("key")
	keys := c.QueryParam("keys")

	userSettings := []models.UserSetting{}

	o := orm.NewOrm()
	qs := o.QueryTable("user_setting")

	qs = qs.Filter("accountId", userId)

	if branchId != "" {
		qs = qs.Filter("branchId", branchId)
	}

	if menuPath != "" {
		qs = qs.Filter("menuPath", menuPath)
	}

	if elementId != "" {
		qs = qs.Filter("elementId", elementId)
	}

	if key != "" {
		qs = qs.Filter("key", key)
	}

	if keys != "" {
		_keys := strings.Split(keys, ",")
		qs = qs.Filter("key__in", _keys)
	}

	_, err := qs.All(&userSettings)
	if err != nil {
		return LoadObjectError(c, err, "USER_SETTING")
	}
	c.JSON(http.StatusOK, userSettings)
	return nil
}
