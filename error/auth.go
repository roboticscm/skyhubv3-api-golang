package error

import (
	"backend/system/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TokenError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusUnauthorized, Message: "Generate Token Error", Key: prefix + ".MSG.GENERATE_TOKEN_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func UnauthorizedError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusUnauthorized, Message: err.Error(), Key: prefix + ".MSG.UNAUTHORIZED_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func CustomHTTPErrorHandler(err error, c echo.Context) {

	if err == middleware.ErrJWTMissing {
		c.Error(echo.NewHTTPError(http.StatusUnauthorized, "Login required"))
		return
	}
	c.Echo().DefaultHTTPErrorHandler(err, c)
}
