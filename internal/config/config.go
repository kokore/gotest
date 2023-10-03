package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	APIConfig   APIConfig
	CovidConfig CovidConfig
}

type APIConfig struct {
	Env     string `envconfig:"ENV" default:"local"`
	BaseURL string `envconfig:"BASE_URL"`
	Port    string `envconfig:"PORT" default:"3000"`
}

type CovidConfig struct {
	Covid_URL string `envconfig:"COVID_API_URL"`
}

func (cfg *AppConfig) Init() {
	envconfig.MustProcess("", &cfg.APIConfig)
	envconfig.MustProcess("", &cfg.CovidConfig)
}

func LoadConfig() *AppConfig {
	_, ok := os.LookupEnv("ENV")
	if !ok {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		err := godotenv.Load(fmt.Sprintf("%v/../../.env.local", basePath))
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}

	cfg := &AppConfig{}
	cfg.Init()
	return cfg
}
