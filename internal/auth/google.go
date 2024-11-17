package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/vincer2040/fuelr/internal/env"
	"github.com/vincer2040/fuelr/internal/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://127.0.0.1:6969/callback-gl",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	oauthStateStringGl = ""
)

func InitializeOAuthGoogle() {
	oauthConfGl.ClientID = env.GetGoogleClientID()
	oauthConfGl.ClientSecret = env.GetGoogleClientSecret()
	oauthStateStringGl = env.GetGoogleOauthStateString()
}

func GoogleLogIn(c echo.Context) error {
	URL, err := url.Parse(oauthConfGl.Endpoint.AuthURL)
	if err != nil {
		return err
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConfGl.ClientID)
	parameters.Add("scope", strings.Join(oauthConfGl.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfGl.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateStringGl)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	c.Response().Header().Add("HX-Redirect", url)
	return nil
}

func GoogleCallBack(c echo.Context) (*types.GoogleUserInfo, error) {
	state := c.FormValue("state")
	if state != oauthStateStringGl {
		c.Response().Header().Add("HX-Redirect", "/")
		return nil, &InvalidState{}
	}

	code := c.FormValue("code")
	if code == "" {
		c.Response().Header().Add("HX-Redirect", "/")
		return nil, &NoCodeProvided{}
	}

	token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var googleUserInfo types.GoogleUserInfo
	err = json.Unmarshal(response, &googleUserInfo)
	if err != nil {
		return nil, err
	}
	return &googleUserInfo, nil
}

type InvalidState struct{}

func (invalidState *InvalidState) Error() string {
	return "Invalid state"
}

type NoCodeProvided struct{}

func (noCodeProvided NoCodeProvided) Error() string {
	return "No code provided"
}
