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
	if res.Origin.Word != "test" {
		t.Fatalf("Wrong original Word")
	}
	if len(res.Origin.Examples) != 8 {
		t.Fatalf("Wrong example counts")
	}
	if len(res.Origin.WordMeanings) != 7 {
		t.Fatalf("Wrong Word meanings counts")
	}
	if res.Translation.Word != "prüfen" {
		t.Fatalf("Wrong Word meanings counts")
	}
	if len(res.Translation.PartOfSpeeches) != 1 {
		t.Fatalf("Wrong Word meanings counts")
	}
}

func TestAllWithDefaultClient(t *testing.T) {
	res, _ := GetTranslatorWithCustomClient(nil).Translate("test", "en", "de")
	if res.Origin.Word != "test" {
		t.Fatalf("Wrong original Word")
	}
	if len(res.Origin.Examples) != 8 {
		t.Fatalf("Wrong example counts")
	}
	if len(res.Origin.WordMeanings) != 7 {
		t.Fatalf("Wrong Word meanings counts")
	}
	if res.Translation.Word != "prüfen" {
		t.Fatalf("Wrong Word meanings counts")
	}
	if len(res.Translation.PartOfSpeeches) != 1 {
		t.Fatalf("Wrong Word meanings counts")
	}
}
