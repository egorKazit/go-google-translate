package server

import "testing"

func TestAllNested(t *testing.T) {
	testArray := []interface{}{[]interface{}{"0"}, []interface{}{"1"}, []interface{}{[]interface{}{"01"}}, []int{1}}

	// one level
	firstLevel := getNested(testArray, []int{0})
	if firstLevel.([]interface{})[0].(string) != "0" {
		t.Fatalf("Can not read one level of array")
	}

	// 2 levels
	secondLevel := getNested(testArray, []int{1, 0})
	if secondLevel != "1" {
		t.Fatalf("Can not read 2 levels of array")
	}

	// out of range
	thirdLevel := getNested(testArray, []int{1, 1})
	if thirdLevel != nil {
		t.Fatalf("Can not handle out of range")
	}

	// non-proper type
	forthLevel := getNested(testArray, []int{3, 0})
	if forthLevel != nil {
		t.Fatalf("Can not handle not interface")
	}

	// more deep level
	fifthLevel := getNested(testArray, []int{7, 0})
	if fifthLevel != nil {
		t.Fatalf("Can not handle not interface")
	}

}

func TestNil(t *testing.T) {
	// one level
	nilValue := getNested(nil, []int{1, 1})
	if nilValue != nil {
		t.Fatalf("Nill expected")
	}
}
