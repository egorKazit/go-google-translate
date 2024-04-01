package server

import (
	"fmt"
	"regexp"
	"testing"
)

func TestPrepareTranslationQuery(t *testing.T) {
	translationObject, _ := prepareTranslationQuery("10101", "01010")
	compiler, _ := regexp.Compile("_reqid=\\d+&bl=01010&f.sid=10101&hl=en-US&id=MkEWBc&rt=c&soc-app=1&soc-device=1&soc-platform=1")
	result := compiler.FindAllString(*translationObject, -1)
	if len(result) != 1 {
		t.Fatalf(fmt.Sprintf(`Incorrect result for word object %s`, result))
	}
}

func TestPrepareTransObject(t *testing.T) {
	transObject, _ := prepareTransObject("test", "en", "ru")
	if *transObject != "[[[\"MkEWBc\",\"[[\\\"test\\\",\\\"en\\\",\\\"ru\\\",true],[null]]\",null,\"generic\"]]]" {
		t.Fatalf(fmt.Sprintf(`Incorrect result for trans object %s`, *transObject))
	}
}
