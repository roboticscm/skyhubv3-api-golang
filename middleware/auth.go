package middleware

import (
	. "backend/system/error"
	. "backend/system/models"
	"backend/system/security"
	"crypto/subtle"

	"github.com/astaxie/beego/orm"
	"github.com/labstack/echo/v4"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	o := orm.NewOrm()
	accounts := []Account{}

	sql := `
		SELECT id, password FROM account
		WHERE disabled = FALSE AND username = ?
	`
	if _, err := o.Raw(sql, username).QueryRows(&accounts); err != nil {
		return false, LoadObjectError(c, err, "ACCOUNT")
	}

	if len(accounts) == 0 {
		return false, Errord400(c, "ACCOUNT.MSG.USERNAME_NOT_FOUND_ERROR", "")
	}

	if len(accounts) > 1 {
		return false, Errord400(c, "ACCOUNT.MSG.TOO_MANY_ACCOUNT_ERROR", "")
	}

	foundEncodedPassword := *accounts[0].Password

	enterEncodedPassword := security.EncodeSHA1Password(password)

	if subtle.ConstantTimeCompare([]byte(foundEncodedPassword), []byte(enterEncodedPassword)) == 1 {
		c.Set("username", username)
		c.Set("userId", accounts[0].Id)
		c.Set("fullName", "//TODO")
		return true, nil
	}
	return false, nil
}
