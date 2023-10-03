package covid

import "github.com/stretchr/testify/mock"

type covidRepositoryMock struct {
	mock.Mock
}

func InitCovidRepositoryMock() *covidRepositoryMock {
	return &covidRepositoryMock{}
}

func (m *covidRepositoryMock) GetCovidCase() (*CovidCase, error) {
	args := m.Called()
	return args.Get(0).(*CovidCase), args.Error(1)
}
