package covid

type CovidCase struct {
	Data []CovidRecord `json:"Data"`
}

type CovidRecord struct {
	ConfirmDate    string      `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            int         `json:"Age"`
	Gender         string      `json:"Gender"`
	GenderEn       string      `json:"GenderEn"`
	Nation         interface{} `json:"Nation"`
	NationEn       string      `json:"NationEn"`
	Province       string      `json:"Province"`
	ProvinceId     int         `json:"ProvinceId"`
	District       interface{} `json:"District"`
	ProvinceEn     string      `json:"ProvinceEn"`
	StatQuarantine int         `json:"StatQuarantine"`
}
