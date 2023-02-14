package config

import (
	"fmt"
	"os"
)

type (
	Resources struct {
		Path string
	}
	Http struct {
		Port string
	}
	Vars struct {
		Http      *Http
		Resources *Resources
	}
	App struct {
		Environment string
		OnDebugMode bool
		IsDEV       bool
		IsPROD      bool
		Vars        *Vars
	}
)

func GetEnvVar(key string, shouldThrowOnMissing bool, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" && shouldThrowOnMissing {
		panic(fmt.Sprintf("Missing environment variable %s", key))
	}
	if value == "" && defaultValue != "" {
		return defaultValue
	}
	return value
}

func GetApp() *App {
	appEnv := GetEnvVar("APP_ENV", false, "DEV")
	debug := GetEnvVar("APP_DEBUG", false, "0")
	return &App{
		Environment: appEnv,
		OnDebugMode: debug == "1",
		IsDEV:       appEnv == "DEV",
		IsPROD:      appEnv == "PROD",
		Vars: &Vars{
			Http: &Http{
				Port: GetEnvVar("APP_PORT", true, ""),
			},
			Resources: &Resources{
				Path: GetEnvVar("RESOURCES_PATH", true, ""),
			},
		},
	}
}
