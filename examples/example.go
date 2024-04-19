package main

import (
	"fmt"
	"os"

	yacloud_translate "github.com/sppps/yacloud-go-translate"
)

func main() {
	tr := yacloud_translate.RestYaTranslate{
		FolderId: os.Getenv("YACLOUD_TRANSLATE_FOLDER_ID"),
		ApiKey:   os.Getenv("YACLOUD_TRANSLATE_API_KEY"),
	}
	result, err := tr.Translate(yacloud_translate.TranslateRequest{
		TargetLanguageCode: "ru",
		Texts:              []string{"hello, world!"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Translations[0].Text)
}
