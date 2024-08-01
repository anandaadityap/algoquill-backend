package configs

import (
	"fmt"
	"os"

	"github.com/anandaadityap/algoquill-backend/internal/pkg/helpers"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT         string
	DATABASE_URL string
}

var ENV *EnvConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		helpers.PanicIfError(err)
	}

	env := EnvConfig{
		PORT:         os.Getenv("PORT"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}

	fmt.Println("Godotenv loaded")

	ENV = &env
}
