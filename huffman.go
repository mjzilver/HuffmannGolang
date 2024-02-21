package main

import "bytes"

const (
	EOF = '□'
)

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
	for len(pq) > 1 {
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

func encode(text string, codes map[rune]string) []byte {
	var buffer bytes.Buffer
	var bits byte
	var bitCount int

	for _, char := range text {
		code := codes[char]
		for _, bit := range code {
			bits = bits << 1
			if bit == '1' {
				bits |= 1
			}
			bitCount++
			if bitCount == 8 {
				buffer.WriteByte(bits)
				bits = 0
				bitCount = 0
			}
		}
	}

	if bitCount > 0 {
		// Pad the remaining bits to form a complete byte
		bits = bits << (8 - bitCount)
		buffer.WriteByte(bits)
	}

	return buffer.Bytes()
}

func decode(encodedText []byte, tree *huffmanNode) string {
	decodedText := ""
	node := tree

	for _, bit := range encodedText {
		for i := 7; i >= 0; i-- {
			if bit&(1<<uint(i)) == 0 {
				node = node.left
			} else {
				node = node.right
			}
			if node.char != 0 {
				char := node.char

				if char == EOF {
					return decodedText
				}

				decodedText += string(char)
				node = tree
			}
		}
	}

	return decodedText
}
