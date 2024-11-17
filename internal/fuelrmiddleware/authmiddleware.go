package fuelrmiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/types"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(types.FuelrContext)
		url := cc.Request().URL.String()
		session, err := cc.Store.Get(cc.Request(), "auth-session")
		if err != nil || session.IsNew {
			if url == "/signin" {
				return next(cc)
			}
			return cc.Redirect(http.StatusSeeOther, "/signin")
		}
		sessionData, err := types.SessionDataFromSession(session)
		if err != nil {
			if url == "/signin" {
				return next(cc)
			}
			return cc.Redirect(http.StatusSeeOther, "/signin")
		}
		if !sessionData.Authenticated {
			if url == "/signin" {
				return next(cc)
			}
			return cc.Redirect(http.StatusSeeOther, "/signin")
		}
		if url == "/signin" {
			return cc.Redirect(http.StatusSeeOther, "/home")
		}
		return next(c)
	}
}
