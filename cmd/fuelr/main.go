package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vincer2040/fuelr/internal/render"
	"github.com/vincer2040/fuelr/internal/routes"
)

func main() {
    e := echo.New()
    e.Renderer = render.New()
    e.Use(middleware.Logger())
    e.Static("/styles", "public/styles")
    e.GET("/", routes.RootGet)
    e.Logger.Fatal(e.Start(":6969"))
}
