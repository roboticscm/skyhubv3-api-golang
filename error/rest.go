package error

import (
	"backend/system/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RestError struct {
	StackTrace string `json:"stackTrace"`
	Message    string `json:"message"`
	Key        string `json:"key"`
	StatusCode int    `json:"statusCode"`
}

func BindObjectError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusBadRequest, Message: "Bind Object Error", Key: prefix + ".MSG.BIND_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func LoadObjectError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusBadRequest, Message: "Load Object Error", Key: prefix + ".MSG.LOAD_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func SaveObjectError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusInternalServerError, Message: "Insert Object Error", Key: prefix + ".MSG.INSERT_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func UpdateObjectError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusInternalServerError, Message: "Update Object Error", Key: prefix + ".MSG.UPDATE_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func DeleteObjectError(c echo.Context, err error, prefix string) error {
	restErr := RestError{StackTrace: err.Error(), StatusCode: http.StatusInternalServerError, Message: "Delete Object Error", Key: prefix + ".MSG.DELETE_OBJECT_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}

func QueryParamError(c echo.Context, message string, prefix string) error {
	restErr := RestError{StackTrace: message, StatusCode: http.StatusBadRequest, Message: message, Key: prefix + ".MSG.QUERY_PARAM_ERROR"}
	slog.Detail(restErr)
	return c.JSON(restErr.StatusCode, restErr)
}
