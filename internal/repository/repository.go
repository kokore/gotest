package repository

import (
	"gotest/internal/config"
	"gotest/internal/repository/covid"
)

type Repositories struct {
	Covid covid.Repo
}

func Init(appConfig *config.AppConfig) Repositories {
	covidRepo := covid.InitCovidRepository(appConfig)

	return Repositories{
		Covid: covidRepo,
	}
}
