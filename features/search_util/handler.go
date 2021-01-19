package search_util

import (
	. "backend/system/error"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func findHandler(c echo.Context) error {
	menuPath := c.QueryParam("menuPath")

	if menuPath == "" {
		return Errord400(c, "SYS.MSG.MISSING_MENU_PATH", "")
	}

	fmt.Println(menuPath)

	// TODO
	c.JSON(http.StatusOK, map[string][]interface{}{
		"fields": {"field1", "field2", "field3"},
	})
	return nil
}
