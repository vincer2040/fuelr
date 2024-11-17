package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/types"
)

func SignOutGet(c echo.Context) error {
	cc := c.(types.FuelrContext)
	session, err := cc.Store.Get(cc.Request(), "auth-session")
	if err != nil {
		return err
	}
	session.Values["authenticated"] = false
	session.Save(cc.Request(), cc.Response())
	cc.Response().Header().Add("HX-Redirect", "/")
	return nil
}
