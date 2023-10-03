package handlers

import (
	"gotest/internal/config"
	"gotest/internal/service"
	"gotest/router/handler/covid"
)

type Handlers struct {
	CovidHandler *covid.CovidHandler
}

func Init(appConfig *config.AppConfig, services service.Service) Handlers {

	return Handlers{
		CovidHandler: covid.Init(appConfig, services.Covid),
	}
}
