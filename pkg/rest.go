package yacloud_translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RestYaTranslate struct {
	FolderId string
	ApiKey   string
	IAMToken string
	Domain   string
	BaseUrl  string
	Logger   *log.Logger
}

type detectLanguageRequest struct {
	DetectLanguageRequest
	FolderId string `json:"folderId"`
}

type listLanguagesRequest struct {
	ListLanguagesRequest
	FolderId string `json:"folderId"`
}

type translateRequest struct {
	TranslateRequest
	FolderId string `json:"folderId"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func stringOrDefault(s1, s2 string) string {
	if len(s1) > 0 {
		return s1
	}
	return s2
}

func (s RestYaTranslate) callRestApi(method string, params any) ([]byte, error) {
	url := fmt.Sprintf("https://%s/%s/%s",
		stringOrDefault(s.Domain, "translate.api.cloud.yandex.net"),
		stringOrDefault(s.BaseUrl, "translate/v2"),
		method)

	if s.Logger != nil {
		s.Logger.Printf("yacloud translate: %s", url)
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	if s.Logger != nil {
		s.Logger.Printf("yacloud translate: %s", string(body))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	if len(s.ApiKey) > 0 {
		req.Header.Set("authorization", fmt.Sprintf("Api-Key %s", s.ApiKey))
	} else if len(s.IAMToken) > 0 {
		req.Header.Set("authorization", fmt.Sprintf("Bearer %s", s.IAMToken))
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if s.Logger != nil {
		s.Logger.Printf("yacloud translate: %s", string(d))
	}

	if resp.StatusCode != http.StatusOK {
		var apiError apiError
		err = json.Unmarshal(d, &apiError)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("api error %d: %s", apiError.Code, apiError.Message)
	}

	return d, err
}

func (s RestYaTranslate) DetectLanguage(req DetectLanguageRequest) (res DetectLanguageResponse, err error) {
	data, err := s.callRestApi("detect", detectLanguageRequest{
		DetectLanguageRequest: req,
		FolderId:              s.FolderId,
	})
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

func (s RestYaTranslate) ListLanguages(req ListLanguagesRequest) (res ListLanguagesResponse, err error) {
	data, err := s.callRestApi("languages", listLanguagesRequest{
		ListLanguagesRequest: req,
		FolderId:             s.FolderId,
	})
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

func (s RestYaTranslate) Translate(req TranslateRequest) (res TranslateResponse, err error) {
	data, err := s.callRestApi("translate", translateRequest{
		TranslateRequest: req,
		FolderId:         s.FolderId,
	})
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}
