package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/auth"
	"github.com/vincer2040/fuelr/internal/db"
	"github.com/vincer2040/fuelr/internal/types"
)

func GoogleAuthGet(c echo.Context) error {
	return auth.GoogleLogIn(c)
}

func GoogleAuthCallBack(c echo.Context) error {
    googleUser, err :=  auth.GoogleCallBack(c)
    if err != nil {
        return err
    }

    cc := c.(types.FuelrContext)
    session, err := cc.Store.Get(cc.Request(), "auth-session")
    if err != nil {
        return err
    }

    user, err := cc.DB.GetUserFromGoogleId(c.Request().Context(), googleUser.GoogleID)
    if err != nil {
        if err != sql.ErrNoRows {
            return err
        }
        newUserId, err := cc.DB.CreateUser(c.Request().Context(), db.CreateUserParams{
            FirstName: googleUser.GivenName,
            LastName: googleUser.FamilyName,
            Email: googleUser.Email,
            Picture: googleUser.Picture,
            AuthMethod: 1,
        })
        if err != nil {
            log.Println("Create User")
            return err
        }
        err = cc.DB.CreateGoogleUser(c.Request().Context(), db.CreateGoogleUserParams{
            GoogleID: googleUser.GoogleID,
            UserID: newUserId,
        })
        if err != nil {
            log.Println("create google user")
            return err
        }
        user = db.User{
            ID: newUserId,
            FirstName: googleUser.GivenName,
            LastName: googleUser.FamilyName,
            Email: googleUser.Email,
            Picture: googleUser.Picture,
        }
    }

    session.Values["authenticated"] = true
    session.Values["userID"] = user.ID
    session.Save(cc.Request(), cc.Response())

    return cc.Redirect(http.StatusSeeOther, "/home")
}
