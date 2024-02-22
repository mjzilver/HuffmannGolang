package main

import (
	"bytes"
	"strings"
)

const (
	SOF = rune(0x02)
	EOF = rune(0x03)
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

	// add custom EOF char to know where to stop
	text += string(EOF)

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
	var sb strings.Builder
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
					return sb.String()
				}

				sb.WriteRune(char)
				node = tree
			}
		}
	}

	return sb.String()
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
		// up the index for the next byte
		*index++
		return &huffmanNode{char: rune(charByte)}
	} else {
		// read left and right
		node := &huffmanNode{}
		node.left = decodeNode(encodedTree, index)
		node.right = decodeNode(encodedTree, index)
		return node
	}
}
