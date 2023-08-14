package config

import (
	"os"

	"github.com/joho/godotenv"
)


var TOKEN_SECRET []byte

func init() {
	godotenv.Load()
	TOKEN_SECRET  = []byte(os.Getenv("TOKEN_SECRET"))
}
