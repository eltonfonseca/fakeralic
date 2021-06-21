package config

import (
	"os"
)

// Config keeps global configs
type Config struct {
	Environment string
	EnvExemple  string
}

var serverConfig *Config

// New retrieves the values from .env and returns a Config.
func New() *Config {
	// Keep only one instance of server config
	if serverConfig != nil {
		return serverConfig
	}

	// Get the current environment
	environment := getEnv("_APP_ENV", "development")

	// If not running on stating nor production, fallback to local configs
	if environment != "staging" && environment != "production" {
		serverConfig = &Config{
			Environment: "development",
			EnvExemple:  "Example of Environment Variable",
		}
		return serverConfig
	}

	// For prod and staging
	serverConfig = &Config{
		Environment: environment,
	}

	return serverConfig
}

func getEnv(k, d string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return d
}
