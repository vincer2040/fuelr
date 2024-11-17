package env

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	return godotenv.Load()
}

func GetGoogleClientID() string {
	return os.Getenv("GOOGLE_CLIENT_ID")
}

func GetGoogleClientSecret() string {
	return os.Getenv("GOOGLE_CLIENT_SECRET")
}

func GetGoogleOauthStateString() string {
	return os.Getenv("GOOGLE_OAUTH_STATE_STRING")
}

func GetStoreSecret() string {
    return os.Getenv("STORE_SECRET")
}
