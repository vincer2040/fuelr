package types

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/db"
)

type FuelrContext struct {
	echo.Context
	Store *sessions.CookieStore
	DB    *db.Queries
}
