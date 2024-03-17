package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	LOG_FILE string
)

func init() {
	godotenv.Load()
	LOG_FILE = os.Getenv("LOG_FILE")
}
