package covid

import (
	"gotest/internal/config"
	response "gotest/internal/error"
	"gotest/internal/service/covid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CovidHandler struct {
	appConfig    *config.AppConfig
	covidService covid.Service
}

func Init(appConfig *config.AppConfig, covidService covid.Service) *CovidHandler {
	return &CovidHandler{
		appConfig:    appConfig,
		covidService: covidService,
	}
}

func (h *CovidHandler) GetCovidSummary(ginCtx *gin.Context) {
	covidSummary, err := h.covidService.GetCovidSummaryService()
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.UnableGetCovidCase, http.StatusBadRequest, err.Error()))
		return
	}

	ginCtx.JSON(http.StatusOK, response.OK(covidSummary))
}
