package error

import (
	"backend/system/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TokenError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusUnauthorized, Message: prefix + ".MSG.GENERATE_TOKEN_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func UnauthorizedError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusUnauthorized, Message: err.Error()}
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
