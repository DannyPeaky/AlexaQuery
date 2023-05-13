package data

type PlayerInfo struct {
	InfoText       InfoText     `json:"infoText"`
	MainArt        MainArt      `json:"mainArt"`
	MediaId        string       `json:"mediaId"`
	MiniInfoText   MiniInfoText `json:"miniInfoText"`
	PlaybackSource string       `json:"playbackSource"`
	Progress       Progress     `json:"progress"`
	Provider       Provider     `json:"provider"`
	State          string       `json:"state"`
}

type InfoText struct {
	Header         string `json:"header"`
	HeaderSubtext1 string `json:"headerSubtext1"`
	MultiLineMode  bool   `json:"multiLineMode"`
	SubText1       string `json:"subText1"`
	SubText2       string `json:"subText2"`
	Title          string `json:"title"`
}

type MainArt struct {
	AltText     string `json:"altText"`
	ContentType string `json:"contentType"`
	ArtType     string `json:"artType"`
	Url         string `json:"url"`
}

type MiniInfoText struct {
	Header         string `json:"header"`
	HeaderSubtext1 string `json:"headerSubtext1"`
	MultiLineMode  bool   `json:"multiLineMode"`
	SubText1       string `json:"subText1"`
	SubText2       string `json:"subText2"`
	Title          string `json:"title"`
}

type Progress struct {
	AllowScrubbing bool   `json:"allowScrubbing"`
	LocationInfo   string `json:"locationInfo"`
	MediaLength    int    `json:"mediaLength"`
	MediaProgress  int    `json:"mediaProgress"`
	ShowTiming     bool   `json:"showTiming"`
	Visible        bool   `json:"visible"`
}

type Provider struct {
	ArtOverlay          string       `json:"artOverlay"`
	FallbackMainArt     string       `json:"fallbackMainArt"`
	ProviderDisplayName string       `json:"providerDisplayName"`
	ProviderLogo        ProviderLogo `json:"providerLogo"`
	ProviderName        string       `json:"providerName"`
}

type ProviderLogo struct {
	AltText     string `json:"altText"`
	ArtType     string `json:"artType"`
	ContentType string `json:"contentType"`
	Url         string `json:"url"`
}
