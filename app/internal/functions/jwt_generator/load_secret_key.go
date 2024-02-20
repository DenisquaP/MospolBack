package generator

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadKey() (key string, err error) {
	err = godotenv.Load(".env")

	if err != nil {
		return
	}

	key = os.Getenv("SECRET_KEY")

	if key == "" {
		return
	}

	return
}
