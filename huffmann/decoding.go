package huffmann

import (
	"bytes"
	"strings"
)

func Decode(encodedText []byte) string {
	var sb strings.Builder
	var tree *huffmanNode

	// Find the index of marker byte
	markerIndex := bytes.IndexByte(encodedText, marker)
	if markerIndex == -1 {
		// No tree marker found, return empty string
		return ""
	}

	tree = decodeTree(encodedText[:markerIndex])

	node := tree

	for _, bit := range encodedText[markerIndex+1:] {
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
