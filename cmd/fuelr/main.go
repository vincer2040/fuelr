package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vincer2040/fuelr/internal/auth"
	"github.com/vincer2040/fuelr/internal/db"
	"github.com/vincer2040/fuelr/internal/env"
	"github.com/vincer2040/fuelr/internal/render"
	"github.com/vincer2040/fuelr/internal/routes"
	"github.com/vincer2040/fuelr/internal/types"
	fuelrsql "github.com/vincer2040/fuelr/sql"
)

func main() {

	ctx := context.Background()

	err := env.InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	auth.InitializeOAuthGoogle()

	storeSecret := env.GetStoreSecret()
	store := sessions.NewCookieStore([]byte(storeSecret))

	dbConn, err := sql.Open("libsql", "file:localdb/local.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbConn.ExecContext(ctx, fuelrsql.DDL)
	if err != nil {
		log.Fatal(err)
	}

	db := db.New(dbConn)

	e := echo.New()
	e.Renderer = render.New()

	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := types.FuelrContext{
				Context: c,
				Store:   store,
				DB:      db,
			}
			return next(cc)
		}
	})

	e.Static("/styles", "public/styles")

	e.GET("/", routes.RootGet)
	e.GET("/signin", routes.SignInGet)

	e.GET("/signin-gl", routes.GoogleAuthGet)
	e.GET("/callback-gl", routes.GoogleAuthCallBack)

	e.Logger.Fatal(e.Start(":6969"))
}
