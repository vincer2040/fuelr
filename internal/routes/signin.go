package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignInGet(c echo.Context) error {
	return c.Render(http.StatusOK, "signin.html", nil)
}
