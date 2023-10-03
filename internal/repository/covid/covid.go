package covid

import (
	"encoding/json"
	"errors"
	"gotest/internal/config"
	"net/http"
)

type repoImpl struct {
	appConfig *config.AppConfig
}

type Repo interface {
	GetCovidCase() (*CovidCase, error)
}

func InitCovidRepository(appConfig *config.AppConfig) Repo {
	return &repoImpl{
		appConfig: appConfig,
	}
}

func (r repoImpl) GetCovidCase() (*CovidCase, error) {
	result, err := http.Get(r.appConfig.CovidConfig.Covid_URL)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, errors.New("can't get covid case")
	}

	var apiResponse CovidCase
	err = json.NewDecoder(result.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
