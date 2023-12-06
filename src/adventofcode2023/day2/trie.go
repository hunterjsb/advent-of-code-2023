package main

import (
	"fmt"
)

var TextDigits = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type node struct {
	children  map[rune]*node
	isWordEnd bool
}

func newNode() *node {
	return &node{
		children: make(map[rune]*node),
	}
}

func GetValidPrefixes(words []string) map[rune]*node {
	var nodes map[rune]*node
	for _, word := range words {
		for _, ch := range word {
			fmt.Printf("%c (%v)", ch, ch)

		}
	}
	return nodes
}
