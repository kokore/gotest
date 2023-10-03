package main

import (
	"gotest/internal/config"
	"gotest/internal/repository"
	"gotest/internal/service"
	"gotest/router"
	handlers "gotest/router/handler"
)

func main() {
	appConfig := config.LoadConfig()

	repositories := repository.Init(appConfig)
	services := service.Init(appConfig, repositories)
	handlers := handlers.Init(appConfig, services)

	router.Init(appConfig, handlers)
}
