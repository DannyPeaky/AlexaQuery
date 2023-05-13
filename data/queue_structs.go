package data

type Queue struct {
	PlayerInfo struct {
		InfoText       *InfoText     `json:"infoText,omitempty"`
		MainArt        *MainArt      `json:"mainArt,omitempty"`
		MediaId        *string       `json:"mediaId,omitempty"`
		MiniInfoText   *MiniInfoText `json:"miniInfoText,omitempty"`
		PlaybackSource *string       `json:"playbackSource,omitempty"`
		Progress       *Progress     `json:"progress,omitempty"`
		Provider       *Provider     `json:"provider,omitempty"`
		State          *string       `json:"state,omitempty"`
	} `json:"playerInfo"`
}

type InfoText struct {
	Header         *string `json:"header,omitempty"`
	HeaderSubtext1 *string `json:"headerSubtext1,omitempty"`
	MultiLineMode  bool    `json:"multiLineMode,omitempty"`
	SubText1       *string `json:"subText1,omitempty"`
	SubText2       *string `json:"subText2,omitempty"`
	Title          *string `json:"title,omitempty"`
}

type MainArt struct {
	AltText     *string `json:"altText,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	ArtType     *string `json:"artType,omitempty"`
	Url         *string `json:"url,omitempty"`
}

type MiniInfoText struct {
	Header         *string `json:"header,omitempty"`
	HeaderSubtext1 *string `json:"headerSubtext1,omitempty"`
	MultiLineMode  bool    `json:"multiLineMode,omitempty"`
	SubText1       *string `json:"subText1,omitempty"`
	SubText2       *string `json:"subText2,omitempty"`
	Title          *string `json:"title,omitempty"`
}

type Progress struct {
	AllowScrubbing bool    `json:"allowScrubbing,omitempty"`
	LocationInfo   *string `json:"locationInfo,omitempty"`
	MediaLength    *int    `json:"mediaLength,omitempty"`
	MediaProgress  *int    `json:"mediaProgress,omitempty"`
	ShowTiming     bool    `json:"showTiming,omitempty"`
	Visible        bool    `json:"visible,omitempty"`
}

type Provider struct {
	ArtOverlay          *string       `json:"artOverlay,omitempty"`
	FallbackMainArt     *string       `json:"fallbackMainArt,omitempty"`
	ProviderDisplayName *string       `json:"providerDisplayName,omitempty"`
	ProviderLogo        *ProviderLogo `json:"providerLogo,omitempty"`
	ProviderName        *string       `json:"providerName,omitempty"`
}

type ProviderLogo struct {
	AltText     *string `json:"altText,omitempty"`
	ArtType     *string `json:"artType,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Url         *string `json:"url,omitempty"`
}
