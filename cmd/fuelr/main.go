package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vincer2040/fuelr/internal/auth"
	"github.com/vincer2040/fuelr/internal/env"
	"github.com/vincer2040/fuelr/internal/render"
	"github.com/vincer2040/fuelr/internal/routes"
)

func main() {
	err := env.InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	auth.InitializeOAuthGoogle()

	e := echo.New()
	e.Renderer = render.New()

	e.Use(middleware.Logger())

	e.Static("/styles", "public/styles")

	e.GET("/", routes.RootGet)
	e.GET("/signin", routes.SignInGet)

	e.GET("/signin-gl", routes.SignInGlGet)
	e.GET("/callback-gl", routes.SignInGoogleCallBack)

	e.Logger.Fatal(e.Start(":6969"))
}
