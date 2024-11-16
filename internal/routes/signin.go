package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/auth"
)

func SignInGet(c echo.Context) error {
	return c.Render(http.StatusOK, "signin.html", nil)
}

func SignInGlGet(c echo.Context) error {
	return auth.GoogleLogIn(c)
}

func SignInGoogleCallBack(c echo.Context) error {
	return auth.GoogleCallBack(c)
}
