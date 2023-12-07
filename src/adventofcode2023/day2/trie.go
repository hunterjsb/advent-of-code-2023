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

func GetPrefixTrie(words []string) map[rune]*node {
	var curNode *node
	rootNodes := make(map[rune]*node)

	for _, word := range words {
		for i, ch := range word {
			//fmt.Printf("%c(%v) ", ch, ch)
			if _, exists := rootNodes[ch]; !exists && i == 0 { // first letter new root node
				rootNodes[ch] = newNode()
				curNode = rootNodes[ch]
			} else if rootNode, exists := rootNodes[ch]; exists && i == 0 { // first letter already has root node
				curNode = rootNode
			} else if _, exists := curNode.children[ch]; !exists && i > 0 { // new child
				curNode.children[ch] = newNode()
				fmt.Printf("appending child %v\n", curNode.children)
			} else if childNode, exists := curNode.children[ch]; exists && i > 0 {
				curNode = childNode
			}
		}
		fmt.Printf("root nodes: %v", rootNodes)
		fmt.Printf("final curnode: %p %v\n", curNode, curNode.children)
		curNode.isWordEnd = true
	}
	return rootNodes
}
