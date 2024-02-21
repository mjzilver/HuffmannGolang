package main

func countFrequency(text string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	return freq
}

func buildHuffmanTree(freq map[rune]int) *huffmanNode {
	var pq PriorityQueue
	for char, freq := range freq {
		node := &huffmanNode{
			char: char,
			freq: freq,
		}

		pq.Push(node)
	}
	for pq.Len() > 1 {
		l1, l2 := pq.Pop(), pq.Pop()

		node := &huffmanNode{
			freq:  l1.freq + l2.freq,
			left:  l1,
			right: l2,
		}
		pq.Push(node)
	}
	// return root
	return pq.Pop()
}

func generateCodes(tree *huffmanNode) map[rune]string {
	codes := make(map[rune]string)
	var walk func(node *huffmanNode, code string)
	walk = func(node *huffmanNode, code string) {
		if node == nil {
			return
		}
		if node.char != 0 {
			codes[node.char] = code
		}
		walk(node.left, code+"0")
		walk(node.right, code+"1")
	}
	walk(tree, "")
	return codes
}

func encode(text string, codes map[rune]string) string {
	encodedText := ""
	for _, char := range text {
		encodedText += codes[char]
	}
	return encodedText
}

func decode(encodedText string, tree *huffmanNode) string {
	decodedText := ""
	node := tree
	for _, bit := range encodedText {
		if bit == '0' {
			node = node.left
		} else {
			node = node.right
		}
		if node.char != 0 {
			decodedText += string(node.char)
			node = tree
		}
	}
	return decodedText
}
