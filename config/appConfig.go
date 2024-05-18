package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerUrl  string
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	// [ENV dev mode] - Only loads the .env file if it is in development mode
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	// [URL] - Recovery URL from application
	httpUrl := os.Getenv("HTTP_URL")

	// [URL][VALIDATION] - Checks if exists value in environment variable
	if len(httpUrl) < 1 {
		return AppConfig{}, errors.New("[ENV] Variable HTTP_URL not found")
	}

	// [PORT] - Recovery PORT from application
	httpPort := os.Getenv("HTTP_PORT")

	// [PORT][VALIDATION] - Checks if exists value in environment variable
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("[ENV] Variable HTTP_PORT not found")
	}

	return AppConfig{
		ServerUrl:  httpUrl,
		ServerPort: httpPort,
	}, nil
}
