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
	var curNode *node
	rootNodes := make(map[rune]*node)

	for _, word := range words {
		for i, ch := range word {
			//fmt.Printf("%c(%v) ", ch, ch)
			if _, exists := rootNodes[ch]; !exists && i == 0 { // first letter is root node
				rootNodes[ch] = newNode()
				curNode = rootNodes[ch]
			}
		}
		fmt.Printf("root nodes: %v", rootNodes)
		fmt.Printf("final curnode: %p %v\n", curNode, curNode.children)
		curNode.isWordEnd = true
	}
	return rootNodes
}
