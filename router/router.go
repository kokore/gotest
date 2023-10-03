package router

import (
	"fmt"
	"gotest/internal/config"
	handlers "gotest/router/handler"

	"github.com/gin-gonic/gin"
)

func Init(appConfig *config.AppConfig, handlers handlers.Handlers) {
	port := fmt.Sprintf(":%s", appConfig.APIConfig.Port)

	engine := gin.Default()

	baseGroup := engine.Group("/api")
	version1 := baseGroup.Group("/v1")

	covidRouters := version1.Group("/covid")
	registerCovidRoutes(covidRouters, handlers)

	engine.Run(port)
}

func registerCovidRoutes(group *gin.RouterGroup, handlers handlers.Handlers) {
	group.GET("/summary", handlers.CovidHandler.GetCovidSummary)
}
