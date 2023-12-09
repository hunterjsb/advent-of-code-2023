package main

import (
	"testing"
)

func TestGetPrefixTrie(t *testing.T) {
	words := TextDigits[:]
	rootNodes := GetPrefixTrie(words)

	// Test that each word is in the trie
	for _, word := range words {
		curNode := rootNodes[rune(word[0])]
		if curNode == nil {
			t.Errorf("Root node for '%c' not found", word[0])
			continue
		}

		for _, ch := range word[1:] {
			curNode = curNode.Traverse(ch)
			if curNode == nil {
				t.Errorf("Node for character '%c' in word '%s' not found in trie", ch, word)
				break
			}
		}

		if !curNode.isWordEnd {
			t.Errorf("Word '%s' not marked as complete in trie", word)
		}
	}
}
