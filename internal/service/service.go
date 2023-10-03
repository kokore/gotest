package service

import (
	"gotest/internal/config"
	"gotest/internal/repository"
	"gotest/internal/service/covid"
)

type Service struct {
	Covid covid.Service
}

func Init(appConfig *config.AppConfig, repos repository.Repositories) Service {
	covidService := covid.InitCovidService(repos.Covid)

	return Service{
		Covid: covidService,
	}
}
