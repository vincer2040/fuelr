package routes

import (

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/auth"
)

func GoogleAuthGet(c echo.Context) error {
	return auth.GoogleLogIn(c)
}

func GoogleAuthCallBack(c echo.Context) error {
    _, err :=  auth.GoogleCallBack(c)
    if err != nil {
        return err
    }
    return nil
}
