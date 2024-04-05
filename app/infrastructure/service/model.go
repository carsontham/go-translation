package service

type TranslateRequest struct {
	Text   string `json:"q"`
	Target string `json:"target"`
}

// TranslateResponse ...
type TranslateResponse struct {
	Data struct {
		Translations struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}
