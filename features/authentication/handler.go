package authentication

import (
	. "backend/system/error"
	"backend/system/gbfunc"
	"backend/system/keys"
	. "backend/system/models"
	"backend/system/security"
	"backend/system/slog"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func loginHandler(c echo.Context) error {
	username := c.Get("username").(string)
	userId := c.Get("userId").(int64)
	fullName := c.Get("fullName").(string)
	accessToken, err := generateToken(false, userId, username, fullName)
	if err != nil {
		return TokenError(c, err, "AUTH")
	}

	refreshToken, err := generateToken(true, userId, username, fullName)
	if err != nil {
		return TokenError(c, err, "AUTH")
	}

	updateFreshToken(c, userId, *refreshToken)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  *accessToken,
		"refreshToken": *refreshToken,
		"userId":       userId,
		"userName":     username,
		"fullName":     fullName,
	})
}

func filter(source []RefreshToken, filterFunc func(RefreshToken) bool) []RefreshToken {
	var dest []RefreshToken
	for _, v := range source {
		if filterFunc(v) == true {
			dest = append(dest, v)
		}
	}
	return dest
}

func logoutHandler(c echo.Context) error {
	userId := c.QueryParam("userId")
	if len(userId) > 0 {
		userId, _ := strconv.ParseInt(userId, 10, 64)
		o := orm.NewOrm()
		sql := `
			DELETE FROM refresh_token
			WHERE account_id = ?
		`
		if _, err := o.Raw(sql, userId).Exec(); err != nil {
			fmt.Println(err)
			return Errord400(c, "SYS.MSG.LOGOUT_ERROR", "")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Logged out",
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Unkown error",
		})
	}

}

func refreshTokenHandler(c echo.Context) error {
	refreshToken := &RefreshToken{}
	if err := c.Bind(refreshToken); err != nil {
		return BindObjectError(c, err, "AUTH")
	}

	if len(*refreshToken.Token) == 0 {
		return UnauthorizedError(c, errors.New("Required Login Error"), "AUTH")
	}

	// check if refresh token existed
	o := orm.NewOrm()
	refreshTokens := []RefreshToken{}

	sql := `
		SELECT id FROM refresh_token
		WHERE token = ?
	`
	if _, err := o.Raw(sql, *refreshToken.Token).QueryRows(&refreshTokens); err != nil {
		return LoadObjectError(c, err, "REFRESH_TOKEN")
	}

	if len(refreshTokens) == 0 {
		return UnauthorizedError(c, errors.New("Required Login Error"), "AUTH")
	}

	j, err := jwt.Parse(*refreshToken.Token, func(token *jwt.Token) (interface{}, error) {
		return keys.VerifyKey, nil
	})

	if err != nil {
		return UnauthorizedError(c, errors.New("Required Login Error"), "AUTH")
	}

	claims := j.Claims.(jwt.MapClaims)
	userId, _ := strconv.ParseInt(claims["userId"].(string), 10, 64)

	newToken, err := generateToken(false, userId, claims["username"].(string), claims["fullName"].(string))
	if err != nil {
		return TokenError(c, err, "AUTH")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":   *newToken,
		"success": true,
	})
}

func generateToken(isRefreshToken bool, userId int64, username string, fullname string) (*string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := t.Claims.(jwt.MapClaims)

	claims["userId"] = strconv.FormatInt(userId, 10)
	claims["username"] = username
	claims["fullName"] = fullname

	const duration = 6000000 //second
	if !isRefreshToken {
		claims["exp"] = time.Now().Add(time.Second * duration).Unix()
	}

	tokenString, err := t.SignedString(keys.SignKey)
	slog.Fatal(err)

	return &tokenString, nil
}

func updateFreshToken(c echo.Context, userId int64, token string) (int64, error) {
	// load old refresh token
	o := orm.NewOrm()
	refreshTokens := []RefreshToken{}

	sql := `
		SELECT * FROM refresh_token
		WHERE account_id = ?
	`
	if _, err := o.Raw(sql, userId).QueryRows(&refreshTokens); err != nil {
		return -1, LoadObjectError(c, err, "REFRESH_TOKEN")
	}

	currentDateTime := gbfunc.GetCurrentMillis()
	if len(refreshTokens) == 0 { // insert refresh token
		refreshToken := RefreshToken{
			Token:     &token,
			AccountId: &userId,
			CreatedAt: &currentDateTime,
		}
		return o.Insert(&refreshToken)
	} else { // update refresh token
		refreshToken := refreshTokens[0]
		refreshToken.Token = &token
		refreshToken.CreatedAt = &currentDateTime
		return o.Update(&refreshToken)
	}
}

func changePasswordHandler(c echo.Context) error {
	userId := gbfunc.GetUserId(c)
	accountInfo := &AccountInfo{}
	if err := c.Bind(accountInfo); err != nil {
		return BindObjectError(c, err, "AUTH")
	}

	o := orm.NewOrm()
	accounts := []Account{}

	sql := `
		SELECT * FROM account
		WHERE id = ?
			AND password = ?
	`
	encodedCurrentPassword := security.EncodeSHA1Password(accountInfo.CurrentPassword)
	if _, err := o.Raw(sql, userId, encodedCurrentPassword).QueryRows(&accounts); err != nil {
		return LoadObjectError(c, err, "ACCOUNT")
	}

	if len(accounts) == 0 {
		return Errord400(c, "SYS.MSG.CURRENT_PASSWORD_IS_INCORRECT", "currentPassword")
	} else {
		account := accounts[0]
		*account.Password = security.EncodeSHA1Password(accountInfo.NewPassword)
		gbfunc.MakeUpdate(c, &account)
		num, err := o.Update(&account)
		if err != nil {
			return Errord400(c, "SYS.MSG.UPDATE_NEW_PASSWORD_ERROR", "")
		} else {
			c.JSON(200, map[string]int64{"messge": num})
		}

		return nil
	}
}

func IsAuthenticated() echo.MiddlewareFunc {
	err := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    keys.VerifyKey,
		SigningMethod: "RS256",
	})

	return err
}
