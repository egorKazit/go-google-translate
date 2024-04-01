package server

import (
	"errors"
	"io"
	"net/http"
	"testing"
)

type CustomReadCloser struct {
	io.ReadCloser
}

func (c CustomReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("")
}

type CustomRoundTripper struct {
	cannotFetchId   bool
	cannotTranslate bool
	isFailedOnFetch bool
}

func (t CustomRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	path := r.URL.Path
	if path == "" && t.cannotFetchId {
		return nil, errors.New("can not fetch id")
	}
	if path == "/_/TranslateWebserverUi/data/batchexecute" && t.cannotTranslate {
		return nil, errors.New("can not fetch id")
	}
	response := &http.Response{}
	if t.isFailedOnFetch {
		response.Body = CustomReadCloser{}
	}
	return response, nil
}

func TestErrorOnIdFetch(t *testing.T) {
	client := http.Client{Transport: CustomRoundTripper{cannotFetchId: true}}
	_, err := GetTranslatorWithCustomClient(&client).Translate("test", "en", "de")
	if err == nil {
		t.Fatalf("Error has to be provided")
	}
}

func TestErrorOnIdFetchRead(t *testing.T) {
	client := http.Client{Transport: CustomRoundTripper{isFailedOnFetch: true}}
	_, err := GetTranslatorWithCustomClient(&client).Translate("test", "en", "de")
	if err == nil {
		t.Fatalf("Error has to be provided")
	}
}

func TestErrorOnTranslate(t *testing.T) {
	client := http.Client{Transport: CustomRoundTripper{cannotTranslate: true}}
	_, err := GetTranslatorWithCustomClient(&client).Translate("test", "en", "de")
	if err == nil {
		t.Fatalf("Error has to be provided")
	}
}

func TestAll(t *testing.T) {
	res, _ := GetTranslator().Translate("test", "en", "de")
	if res.origin.word != "test" {
		t.Fatalf("Wrong original word")
	}
	if len(res.origin.examples) != 8 {
		t.Fatalf("Wrong example counts")
	}
	if len(res.origin.wordMeanings) != 7 {
		t.Fatalf("Wrong word meanings counts")
	}
	if res.translation.word != "prüfen" {
		t.Fatalf("Wrong word meanings counts")
	}
	if len(res.translation.partOfSpeeches) != 1 {
		t.Fatalf("Wrong word meanings counts")
	}
}

func TestAllWithDefaultClient(t *testing.T) {
	res, _ := GetTranslatorWithCustomClient(nil).Translate("test", "en", "de")
	if res.origin.word != "test" {
		t.Fatalf("Wrong original word")
	}
	if len(res.origin.examples) != 8 {
		t.Fatalf("Wrong example counts")
	}
	if len(res.origin.wordMeanings) != 7 {
		t.Fatalf("Wrong word meanings counts")
	}
	if res.translation.word != "prüfen" {
		t.Fatalf("Wrong word meanings counts")
	}
	if len(res.translation.partOfSpeeches) != 1 {
		t.Fatalf("Wrong word meanings counts")
	}
}
