package error

import (
	"backend/system/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomError struct {
	StackTrace string `json:"stackTrace"`
	Message    string `json:"message"`
	Field      string `json:"field"`
	StatusCode int    `json:"statusCode"`
}

func BindObjectError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusBadRequest, Message: prefix + ".MSG.BIND_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func LoadObjectError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusBadRequest, Message: prefix + ".MSG.LOAD_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func SaveObjectError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusInternalServerError, Message: prefix + ".MSG.INSERT_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func UpdateObjectError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusInternalServerError, Message: prefix + ".MSG.UPDATE_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func DeleteObjectError(c echo.Context, err error, prefix string) error {
	restErr := CustomError{StackTrace: err.Error(), StatusCode: http.StatusInternalServerError, Message: prefix + ".MSG.DELETE_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func QueryParamError(c echo.Context, message string, prefix string) error {
	restErr := CustomError{StackTrace: message, StatusCode: http.StatusBadRequest, Message: prefix + ".MSG.QUERY_PARAM_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}
