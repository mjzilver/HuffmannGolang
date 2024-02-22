package huffmann

func buildHuffmanTree(freq map[rune]int) *huffmanNode {
	var pq PriorityQueue
	for char, freq := range freq {
		node := &huffmanNode{
			char: char,
			freq: freq,
		}

		pq.Push(node)
	}
	for len(pq) > 1 {
		l1, l2 := pq.Pop(), pq.Pop()

		node := &huffmanNode{
			freq:  l1.freq + l2.freq,
			left:  l1,
			right: l2,
		}
		pq.Push(node)
	}
	return pq.Pop()
}

func encodeTree(tree *huffmanNode) []byte {
	buffer := []byte{}
	encodeNode(tree, &buffer)
	return buffer
}

func encodeNode(node *huffmanNode, buffer *[]byte) {
	if node.IsLeaf() {
		*buffer = append(*buffer, 1)
		*buffer = append(*buffer, byte(node.char))
	} else {
		*buffer = append(*buffer, 0)
		encodeNode(node.left, buffer)
		encodeNode(node.right, buffer)
	}
}

func decodeTree(encodedTree []byte) *huffmanNode {
	var index int
	return decodeNode(encodedTree, &index)
}

func decodeNode(encodedTree []byte, index *int) *huffmanNode {
	byte := encodedTree[*index]
	*index++
	// leaf is 1
	if byte == 1 {
		charByte := encodedTree[*index]
		*index++
		return &huffmanNode{char: rune(charByte)}
	} else {
		node := &huffmanNode{}
		node.left = decodeNode(encodedTree, index)
		node.right = decodeNode(encodedTree, index)
		return node
	}
}
