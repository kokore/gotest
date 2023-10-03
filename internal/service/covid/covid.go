package covid

import (
	"gotest/internal/config"
	"gotest/internal/repository/covid"
)

type serviceImpl struct {
	appConfig       *config.AppConfig
	covidRepository covid.Repo
}

type Service interface {
	GetCovidSummaryService() (*CovidSummary, error)
}

func InitCovidService(appConfig *config.AppConfig, covidRepository covid.Repo) Service {
	return &serviceImpl{
		appConfig:       appConfig,
		covidRepository: covidRepository,
	}
}

func (s serviceImpl) GetCovidSummaryService() (*CovidSummary, error) {

	covidCase, err := s.covidRepository.GetCovidCase()
	if err != nil {
		return nil, err
	}

	provinceCount := make(map[string]int)
	ageGroupCount := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	for _, record := range covidCase.Data {
		provinceCount[record.Province]++
		ageGroup := getAgeGroup(record.Age)
		ageGroupCount[ageGroup]++
	}

	covidSum := &CovidSummary{
		Province: provinceCount,
		AgeGroup: ageGroupCount,
	}

	return covidSum, nil
}

func getAgeGroup(age int) string {
	switch {
	case age >= 0 && age <= 30:
		return "0-30"
	case age >= 31 && age <= 60:
		return "31-60"
	case age >= 61:
		return "61+"
	default:
		return "N/A"
	}
}
