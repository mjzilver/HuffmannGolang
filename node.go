package main

type huffmanNode struct {
	char  rune
	freq  int
	left  *huffmanNode
	right *huffmanNode
}

// func to print huffmanNode
func (node *huffmanNode) String() string {
	var left, right string
	if node.left != nil {
		left = node.left.String()
	}
	if node.right != nil {
		right = node.right.String()
	}
	return left + string(node.char) + right
}
