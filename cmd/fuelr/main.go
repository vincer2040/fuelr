package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fueler/internal/env"
)

func main() {
    handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{})
    slog.SetDefault(slog.New(handler))

	err := env.Init()
	if err != nil {
        slog.Error("failed to initialize .env", slog.Any("err", err))
		return
	}

	port := env.GetPort()

	e := echo.New()

    e.Start(":" + port)
}
