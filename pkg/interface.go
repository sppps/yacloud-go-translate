package yacloud_translate

type YaTranslate interface {
	DetectLanguage(req DetectLanguageRequest) (DetectLanguageResponse, error)
	ListLanguages(req ListLanguagesRequest) (ListLanguagesResponse, error)
	Translate(req TranslateRequest) (TranslateResponse, error)
}

type DetectLanguageRequest struct {
	Text              string   `json:"text"`
	LanguageCodeHints []string `json:"languageCodeHints"`
}

type DetectLanguageResponse struct {
	LanguageCode string `json:"languageCode"`
}

type ListLanguagesRequest struct {
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ListLanguagesResponse struct {
	Languages []Language `json:"languages"`
}

type Format string

const (
	FormatPlainText Format = "PLAIN_TEXT"
	FormatHtml      Format = "HTML"
)

type TranslateGlossaryConfig struct {
	GlossaryData GlossaryData `json:"glossaryData"`
}

type GlossaryData struct {
	GlossaryPairs []GlossaryPair `json:"glossaryPairs"`
}

type GlossaryPair struct {
	SourceText     string `json:"sourceText"`
	TranslatedText string `json:"translatedText"`
	Exact          bool   `json:"exact,omitempty"`
}

type TranslateRequest struct {
	SourceLanguageCode string                   `json:"sourceLanguageCode,omitempty"`
	TargetLanguageCode string                   `json:"targetLanguageCode"`
	Format             Format                   `json:"format,omitempty"`
	Texts              []string                 `json:"texts"`
	Model              string                   `json:"model,omitempty"`
	GlossaryConfig     *TranslateGlossaryConfig `json:"glossaryConfig,omitempty"`
	Speller            bool                     `json:"speller,omitempty"`
}

type TranslateResponse struct {
	Translations []TranslatedText
}

type TranslatedText struct {
	Text                 string
	DetectedLanguageCode string
}
