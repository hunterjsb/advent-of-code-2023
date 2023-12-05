package main

import "fmt"

var textDigits = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func GetValidPrefixes(words *[]string) []string {
	prefixes := make([]string, len(*words)) // since it will be at least words long if all word in words are unique
	for _, ch := range *words {
		fmt.Println(ch)
	}
	return prefixes
}
