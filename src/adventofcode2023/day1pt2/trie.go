package main

type Node struct {
	children  map[rune]*Node
	isWordEnd bool
}

func newNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
	}
}

func GetPrefixTrie(words []string) map[rune]*Node {
	rootNodes := make(map[rune]*Node)

	for _, word := range words {
		var curNode *Node

		for i, ch := range word {
			if i == 0 {
				if _, exists := rootNodes[ch]; !exists {
					rootNodes[ch] = newNode() // Create a new root node if it doesn't exist
				}
				curNode = rootNodes[ch]
			} else {
				if _, exists := curNode.children[ch]; !exists {
					curNode.children[ch] = newNode() // Create a new node if it doesn't exist
				}
				curNode = curNode.children[ch] // Move to the child node
			}
		}
		curNode.isWordEnd = true // Mark the end of the word
	}
	return rootNodes
}

func (n *Node) Traverse(ch rune) *Node {
	if child, exists := n.children[ch]; exists {
		return child // Return the child node
	}
	return nil // Return nil if the child node does not exist
}
