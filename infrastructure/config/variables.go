package config

import (
	"fmt"
	"os"
)

type (
	httpVars struct {
		Port string
	}
	vars struct {
		IsDev  bool
		IsProd bool
		HTTP   *httpVars
	}
)

func getEnv(key string, shouldThrowOnMissing bool, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" && shouldThrowOnMissing {
		panic(fmt.Sprintf("Missing environment variable %s", key))
	}
	if value == "" && defaultValue != "" {
		return defaultValue
	}
	return value
}

func GetEnvVars() *vars {
	appEnv := getEnv("APP_ENV", false, "DEV")
	return &vars{
		IsDev:  appEnv == "ENV",
		IsProd: appEnv == "PROD",
		HTTP: &httpVars{
			Port: getEnv("APP_PORT", false, "8080"),
		},
	}
}
