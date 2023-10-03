package covid

type CovidSummary struct {
	Province map[string]int `json:"province"`
	AgeGroup map[string]int `json:"ageGroup"`
}
