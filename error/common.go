package error

import (
	"backend/system/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CommonError(c echo.Context, message string, key string) error {
	restErr := RestError{StackTrace: message, StatusCode: http.StatusBadRequest, Message: message, Key: key}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}
