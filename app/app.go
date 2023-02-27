package app

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
	Vars struct {
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
		DB          *JsonDB
		Services    *Services
	}
)

var app *App = nil

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
	// keeps app as singleton unique per runtime (pointer)
	if app != nil {
		return app
	}

	// envs
	appEnv := GetEnvVar("ENV_MODE", false, "DEV")
	debug := GetEnvVar("DEBUG", false, "1")

	// paths
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting pwd: %s , defining as empty string", err)
		pwd = ""
	}

	resourcesPath := path.Join(pwd, "resources")
	databasePath := path.Join(resourcesPath, "database")
	yamlPath := path.Join(resourcesPath, "yaml")

	app = &App{
		Environment: appEnv,
		OnDebugMode: debug == "1",
		IsDEV:       appEnv == "DEV",
		IsPROD:      appEnv == "PROD",
		Vars: &Vars{
			Resources: &Resources{
				Path: GetEnvVar("RESOURCES_PATH", false, "resources"),
			},
		},
		Paths: &Paths{
			PWD:       pwd,
			Resources: resourcesPath,
			Database:  databasePath,
			Yaml:      yamlPath,
		},
		DB:       nil,
		Services: nil,
	}

	// Initializing json db
	app.InitJsonDB()

	app.Services = &Services{
		db:   app.DB,
		vars: app.Vars,
	}

	return app
}
