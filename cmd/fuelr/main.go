package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fueler/internal/env"
<<<<<<< HEAD
)

func main() {
    handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{})
    slog.SetDefault(slog.New(handler))

	err := env.Init()
	if err != nil {
        slog.Error("failed to initialize .env", slog.Any("err", err))
=======
	"github.com/vincer2040/fueler/internal/render"
	"github.com/vincer2040/fueler/internal/routes"
)

func main() {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{})
	slog.SetDefault(slog.New(handler))

	err := env.Init()
	if err != nil {
		slog.Error("failed to initialize .env", slog.Any("err", err))
>>>>>>> c13695c (batman)
		return
	}

	port := env.GetPort()

	e := echo.New()

<<<<<<< HEAD
    e.Start(":" + port)
=======
	e.Renderer = render.New()
	e.Static("styles/", "public/styles/")

	e.GET("/", routes.RootGet)

	e.Start(":" + port)
>>>>>>> c13695c (batman)
}
