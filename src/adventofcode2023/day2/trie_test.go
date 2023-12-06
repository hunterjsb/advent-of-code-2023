package main

import (
	"reflect"
	"testing"
)

func TestGetValidPrefixes(t *testing.T) {
	// testWords := []string{"zero", "one", "two"}
	// expectedPrefixes := []string{"zero", "one", "two"}

	actualPrefixes := GetValidPrefixes(TextDigits[:])

	if !reflect.DeepEqual(actualPrefixes, TextDigits) {
		t.Errorf("getValidPrefixes() = %v; want %v", actualPrefixes, TextDigits)
	}
}
