package data

type Queue struct {
	PlayerInfo PlayerInfo `json:"playerInfo"`
}

type PlayerInfo struct {
	InfoText struct {
		SubText1 string `json:"subText1"`
		Title    string `json:"title"`
	} `json:"infoText"`

	MainArt struct {
		Url string `json:"url"`
	} `json:"mainArt"`

	Progress struct {
		MediaLength   int `json:"mediaLength"`
		MediaProgress int `json:"mediaProgress"`
	} `json:"progress"`

	Provider struct {
		ProviderName string `json:"providerName"`
	} `json:"provider"`

	State string `json:"state"`
}
