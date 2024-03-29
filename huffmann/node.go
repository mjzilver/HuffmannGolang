package huffmann

type huffmanNode struct {
	char  rune
	freq  int
	left  *huffmanNode
	right *huffmanNode
}

// debug function to print huffmanNode
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

func (node *huffmanNode) IsLeaf() bool {
	return node.left == nil && node.right == nil
}
