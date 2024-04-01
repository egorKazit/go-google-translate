package server

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testStructure2Extreact struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

func TestExtractorMarshalledJson(t *testing.T) {
	testData := testStructure2Extreact{"firstValue", "secondValue"}
	testDataInString, _ := json.Marshal(testData)
	firstValue := extractKey("field1", string(testDataInString))
	if firstValue != "firstValue" {
		t.Fatalf(fmt.Sprintf(`Can not extract value %s`, firstValue))
	}
	secondValue := extractKey("field2", string(testDataInString))
	if secondValue != "secondValue" {
		t.Fatalf(fmt.Sprintf(`Can not extract value %s`, firstValue))
	}
}

func TestExtractorJsonInMiddle(t *testing.T) {
	testData := testStructure2Extreact{"firstValue", "secondValue"}
	testDataInString, _ := json.Marshal(testData)
	firstValue := extractKey("field1", fmt.Sprintf(`test %s test`, testDataInString))
	if firstValue != "firstValue" {
		t.Fatalf(fmt.Sprintf(`Can not extract value %s`, firstValue))
	}
}

func TestExtractorNoValue(t *testing.T) {
	firstValue := extractKey("field1", `test test`)
	if firstValue != "" {
		t.Fatalf(fmt.Sprintf(`Can not extract value %s`, firstValue))
	}
}

func TestExtractorErrorValue(t *testing.T) {
	firstValue := extractKey("[]", `test test`)
	if firstValue != "" {
		t.Fatalf(fmt.Sprintf(`Can not extract value %s`, firstValue))
	}
}
