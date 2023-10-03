package covid

type CovidCase struct {
	Data []CovidRecord `json:"Data"`
}

type CovidRecord struct {
	Age      int    `json:"Age"`
	Province string `json:"Province"`
}
