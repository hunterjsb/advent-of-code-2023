package main

import (
	"reflect"
	"testing"
)

func TestGetValidPrefixes(t *testing.T) {
	testWords := []string{"zero", "one", "two"}
	expectedPrefixes := []string{"zero", "one", "two"}

	actualPrefixes := GetValidPrefixes(testWords)

	if !reflect.DeepEqual(actualPrefixes, expectedPrefixes) {
		t.Errorf("getValidPrefixes() = %v; want %v", actualPrefixes, expectedPrefixes)
	}
}
