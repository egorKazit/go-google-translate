package server

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"math/rand"
)

type urlData struct {
	CID       string `url:"id"`
	SID       string `url:"f.sid"`
	Bl        string `url:"bl"`
	Hl        string `url:"hl"`
	App       int    `url:"soc-app"`
	Platform  int    `url:"soc-platform"`
	Device    int    `url:"soc-device"`
	RequestID int    `url:"_reqid"`
	Rt        string `url:"rt"`
}

func prepareTranslationQuery(sid string, bl string) (*string, error) {
	// create url object
	urlDataObject := &urlData{"MkEWBc", sid, bl, "en-US", 1, 1, 1, rand.Intn(90000), "c"}
	// prepare query values
	urlDataObjectQuery, _ := query.Values(urlDataObject)
	// encode at last
	result := urlDataObjectQuery.Encode()
	return &result, nil
}

func prepareTransObject(text string, sourceLang string, targetLang string) (*string, error) {
	// create object array and marshal + handle error if so
	translationArray := []any{[]any{text, sourceLang, targetLang, true}, []any{nil}}
	translationArrayBytes, _ := json.Marshal(translationArray)
	// create final object and marshal + handle error if so
	transArray := []any{[]any{[]any{"MkEWBc", string(translationArrayBytes), nil, "generic"}}}
	transArrayResult, _ := json.Marshal(transArray)
	// covert to string and return
	result := string(transArrayResult)
	return &result, nil
}
