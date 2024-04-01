package server

import (
	"os"
	"testing"
)

func TestResolveWrong(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveNoJson(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n\"test\""))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveWrongJsonEmptyObject(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n{}"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveWrongJsonEmptyArray(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n[]"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveWrongJsonShortArray(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n[[],[],[],[]]"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveWrongJsonEmptyValue(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n[[[]],[[]],[[]]]"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveWrongJsonNoNestedArray(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n[{},{},{},{}]"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolveWrongJsonNoNestedMarshaling(t *testing.T) {
	res, err := resolveTranslation([]byte(")]}'\n\n7156\n[[\"\",\"\",\"tt\"]]"))
	if res != nil {
		t.Fatalf("Incorect result")
	}
	if err == nil {
		t.Fatalf("Error not provided")
	}
}

func TestResolve(t *testing.T) {
	data, err := os.ReadFile("../../resources/response.txt")
	if err != nil {
		t.Fatalf("Error provided")
	}
	res, err := resolveTranslation(data)
	if res == nil {
		t.Fatalf("Incorect result")
	}
	if len(res.OtherTranslations) != 19 {
		t.Fatalf("Can not parse")
	}
	if err != nil {
		t.Fatalf("Error provided")
	}
}

func TestGetOriginalWord(t *testing.T) {
	originWord := getOriginalWord(nil)
	if originWord.Word != "" {
		t.Fatalf("Origin should be initial: Word")
	}
	if originWord.Examples != nil {
		t.Fatalf("Origin should be initial: Examples")
	}
	if originWord.WordMeanings != nil {
		t.Fatalf("Origin should be initial: meanings")
	}
}

func TestGetOtherTranslations(t *testing.T) {
	otherTranslations := getOtherTranslations(nil)
	if len(otherTranslations) != 0 {
		t.Fatalf("Other translations should be initial")
	}
}

func TestGetPartOfSpeech(t *testing.T) {
	partsOfSpeech := getPartOfSpeech(nil)
	if len(partsOfSpeech) != 0 {
		t.Fatalf("Other translations should be initial")
	}
}

func TestGetNilSynonyms(t *testing.T) {
	synonyms := getSynonyms(nil)
	if len(synonyms) != 0 {
		t.Fatalf("Synonyms should be initial")
	}
}

func TestGetNotNilSynonyms(t *testing.T) {
	synonyms := getSynonyms([]interface{}{nil})
	if len(synonyms) != 0 {
		t.Fatalf("Synonyms should be initial")
	}
}

func TestGetNilExamples(t *testing.T) {
	examples := getExamples(nil)
	if len(examples) != 0 {
		t.Fatalf("Examples should be initial")
	}
}

func TestGetNotNilExamples(t *testing.T) {
	examples := getExamples([]interface{}{nil})
	if len(examples) != 0 {
		t.Fatalf("Examples should be initial")
	}
}

func TestGetNilMeanings(t *testing.T) {
	meanings := getMeanings(nil)
	if len(meanings) != 0 {
		t.Fatalf("Meanings should be initial")
	}
}

func TestGetNilMeanings1Level(t *testing.T) {
	meanings := getMeanings([]interface{}{nil})
	if len(meanings) != 0 {
		t.Fatalf("Meanings should be initial")
	}
}
func TestGetNilMeanings2Levels(t *testing.T) {
	meanings := getMeanings([]interface{}{[]interface{}{nil, nil}})
	if len(meanings) != 0 {
		t.Fatalf("Meanings should be initial")
	}
}

func TestGetNilMeanings3Levels(t *testing.T) {
	meanings := getMeanings([]interface{}{[]interface{}{[]interface{}{nil}, []interface{}{nil}}})
	if len(meanings) != 0 {
		t.Fatalf("Meanings should be initial")
	}
}
