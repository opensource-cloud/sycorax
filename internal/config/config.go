package config

import (
	"fmt"
	jsondb "github.com/opensource-cloud/sycorax/internal/database/engine"
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
	Config struct {
		Environment string
		OnDebugMode bool
		IsDEV       bool
		IsPROD      bool
		Vars        *Vars
		Paths       *Paths
		DB          *jsondb.JsonDB
	}
)

// GetEnvVar Returns an env var by key, can throw a error and return a default value as string
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

var config *Config = nil

// NewConfig Returns a new Config struct
func NewConfig() *Config {
	// keeps api as singleton unique per runtime (pointer)
	if config != nil {
		return config
	}

	// envs
	appEnv := GetEnvVar("ENV_MODE", false, "DEV")
	debug := GetEnvVar("DEBUG", false, "1")

	// paths
	pwd, err := os.Getwd()
	if err != nil {
		pwd = ""
	}

	resourcesPath := path.Join(pwd, "resources")
	databasePath := path.Join(resourcesPath, "database")
	yamlPath := path.Join(resourcesPath, "yaml")

	db, err := jsondb.NewJsonDB(&jsondb.JsonDBConfig{
		Path:   databasePath,
		DBName: "sycorax.db",
	})
	if err != nil {
		panic(err)
	}

	config = &Config{
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
		DB: db,
	}

	return config
}
