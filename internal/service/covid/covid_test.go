package covid_test

import (
	"gotest/internal/repository/covid"
	covidService "gotest/internal/service/covid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCovidSummaryService(t *testing.T) {

	type testCase struct {
		name     string
		age      int
		province string
		expected covidService.CovidSummary
	}

	cases := []testCase{
		{name: "age 0", age: 0, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 1, "31-60": 0, "61+": 0, "N/A": 0},
		}},
		{name: "age 30", age: 30, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 1, "31-60": 0, "61+": 0, "N/A": 0},
		}},
		{name: "age 31", age: 31, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 0, "31-60": 1, "61+": 0, "N/A": 0},
		}},
		{name: "age 60", age: 60, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 0, "31-60": 1, "61+": 0, "N/A": 0},
		}},
		{name: "age 61", age: 61, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 0, "31-60": 0, "61+": 1, "N/A": 0},
		}},
		{name: "age 100", age: 100, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 0, "31-60": 0, "61+": 1, "N/A": 0},
		}},
		{name: "age -1", age: -1, province: "a", expected: covidService.CovidSummary{
			Province: map[string]int{"a": 1},
			AgeGroup: map[string]int{"0-30": 0, "31-60": 0, "61+": 0, "N/A": 1},
		}},
	}

	for _, tc := range cases {

		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			covidRepo := covid.InitCovidRepositoryMock()
			data := covid.CovidCase{
				Data: []covid.CovidRecord{
					{
						Age:      tc.age,
						Province: tc.province,
					},
				},
			}
			covidRepo.On("GetCovidCase").Return(&data, nil)
			covidService := covidService.InitCovidService(covidRepo)

			// Act
			summary, _ := covidService.GetCovidSummaryService()
			expected := tc.expected

			// Assert
			assert.Equal(t, &expected, summary)
		})
	}
}
