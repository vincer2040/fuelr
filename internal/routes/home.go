package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeGet(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"FirstName": "",
	})
}
