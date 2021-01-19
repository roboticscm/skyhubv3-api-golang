package error

import (
	"backend/system/slog"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Errord400(c echo.Context, message string, field string) error {
	restErr := CustomError{StackTrace: message, StatusCode: http.StatusBadRequest, Message: message, Field: field}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func ExistedError(c echo.Context, field string) error {
	restErr := CustomError{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("SYS.MSG.%v_IS_EXISTED", strings.ToUpper(field)), Field: field}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}


func DuplicatedError(c echo.Context, field string) error {
	restErr := CustomError{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("SYS.MSG.%v_IS_DUPLICATED", strings.ToUpper(field)), Field: field}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}