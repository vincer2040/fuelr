package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/types"
)

func MeGet(c echo.Context) error {
	cc := c.(types.FuelrContext)
	session, err := cc.Store.Get(cc.Request(), "auth-session")
	if err != nil {
		return err
	}
	sessionData, err := types.SessionDataFromSession(session)
	if err != nil {
		return err
	}
	user, err := cc.DB.GetUserById(context.TODO(), sessionData.UserID)
	if err != nil {
		return err
	}
	return cc.Render(http.StatusOK, "me.html", map[string]interface{}{
		"Route":   "Me",
		"Picture": user.Picture,
	})
}
