package config

import (
	"os"
	"strings"

	"github.com/bxxf/go-template/internal/logger"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var ENV_VALUES = []string{"PORT"}

type Config struct {
	Port string
}

func NewConfig(lc fx.Lifecycle, logger *logger.LoggerInstance) *Config {
	values := load(logger)

	return &Config{
		Port: values["PORT"],
	}
}

func load(logger *logger.LoggerInstance) map[string]string {
	// Create map for environment variables
	var values map[string]string = make(map[string]string)
	// Load .env file only if not in production
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			logger.Error("Error loading .env file")
		}
	}
	// Loop through environment variables and check if they exist
	for _, key := range ENV_VALUES {
		values[key] = os.Getenv(key)
		if values[key] == "" {
			// if the name of the key includes optional then it is not required
			if !strings.Contains(strings.ToLower(key), "optional") {
				logger.Error("Missing environment variable: " + key)

				os.Exit(1)
			} else {
				logger.Warn("Optional environment variable not found: " + key)
			}

		}
	}

	return values
}
