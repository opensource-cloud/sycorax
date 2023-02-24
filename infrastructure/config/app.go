package config

import (
	"fmt"
	"log"
	"os"
	"path"
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
	Paths struct {
		PWD       string
		Resources string
		Database  string
		Yaml      string
	}
	App struct {
		Environment string
		OnDebugMode bool
		IsDEV       bool
		IsPROD      bool
		Vars        *Vars
		Paths       *Paths
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
	// envs
	appEnv := GetEnvVar("APP_ENV", false, "DEV")
	debug := GetEnvVar("APP_DEBUG", false, "0")

	// paths
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting pwd: %s , defining as empty string", err)
		pwd = ""
	}
	resourcesPath := path.Join(pwd, "resources")
	databasePath := path.Join(resourcesPath, "database")
	yamlPath := path.Join(resourcesPath, "yaml")

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
		Paths: &Paths{
			PWD:       pwd,
			Resources: resourcesPath,
			Database:  databasePath,
			Yaml:      yamlPath,
		},
	}
}
