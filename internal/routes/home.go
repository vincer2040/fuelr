package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/types"
)

func HomeGet(c echo.Context) error {
    cc := c.(types.FuelrContext)
    session, err := cc.Store.Get(cc.Request(), "auth-session")
    if err != nil {
        return err
    }

    sessionData, err := types.SessionDataFromSession(session)
    if err != nil {
        return err
    }

    firstName, err := cc.DB.GetFirstNameById(context.TODO(), sessionData.UserID)

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"FirstName": firstName,
	})
}
