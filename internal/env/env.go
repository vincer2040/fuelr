package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() error {
    return godotenv.Load()
}

func GetPort() string {
    return os.Getenv("PORT")
}
