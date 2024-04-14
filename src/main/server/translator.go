package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const googleUrl = "https://translate.google.com"
const googleBatchPath = "/_/TranslateWebserverUi/data/batchexecute?"

func GetTranslator() *translator {
	return &translator{&http.Client{}}
}

func GetTranslatorWithCustomClient(client *http.Client) *translator {
	if client != nil {
		return &translator{client}
	} else {
		log.Print("provided client is nil, so default will be used")
		return &translator{&http.Client{}}
	}
}

type translator struct {
	client *http.Client
}

func (translator *translator) Translate(text string, sourceLang string, targetLang string) (*Translation, error) {

	urlData, err := translator.fetchURLData()
	if err != nil {
		return nil, err
	}

	translation, err := translator.fetchTranslation(text, sourceLang, targetLang, *urlData)
	if err != nil {
		return nil, err
	}
	return translation, nil
}

func (translator *translator) fetchURLData() (result *string, err error) {
	// prepare new request to translator
	request, err := http.NewRequest(http.MethodGet, googleUrl, nil)
	if err != nil {
		log.Printf("Error at fetch SID URL creation: %s", err.Error())
		return nil, err
	}

	// execute request
	response, err := translator.client.Do(request)
	if err != nil {
		log.Printf("Error at fetch SID: %s", err.Error())
		return nil, err
	}

	// handle response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error at read response: %s", err.Error())
		return nil, err
	}
	bodyString := string(body)

	// prepare URL data
	return prepareTranslationQuery(extractKey("FdrFJe", bodyString), extractKey("cfb2h", bodyString))

}

func (translator *translator) fetchTranslation(text string, sourceLang string, targetLang string, translationQuery string) (*Translation, error) {

	// prepare body
	data := url.Values{}
	res, err := prepareTransObject(text, sourceLang, targetLang)
	if err != nil {
		log.Printf("Error at preparation: %s", err.Error())
		return nil, err
	}
	data.Set("f.req", *res)

	// create request
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s%s", googleUrl, googleBatchPath, translationQuery),
		strings.NewReader(data.Encode()))
	// handle error if so
	if err != nil {
		log.Printf("Error at request preparation: %s", err.Error())
		return nil, err
	}
	request.Header.Set("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	response, err := translator.client.Do(request)
	if err != nil {
		log.Printf("Error at request send: %s", err.Error())
		return nil, err
	}
	body, _ := io.ReadAll(response.Body)
	translation, err := resolveTranslation(text, body)
	if err != nil {
		log.Printf("Error at Word resolving: %s", err.Error())
		return nil, err
	}
	return translation, nil
}
