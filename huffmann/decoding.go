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

	// last byte is uint8 for bits that are padding
	padding := int(encodedText[len(encodedText)-1])
	// Remove everthing before (and including) marker and padding byte
	relevantBytes := encodedText[markerIndex+1 : len(encodedText)-1]

	for index, bit := range relevantBytes {
		// Determine the number of padding bits
		paddingBits := 0
		if index == len(relevantBytes)-1 {
			paddingBits = padding
		}

		// Iterate over the bits, excluding padding bits
		for i := 7; i >= paddingBits; i-- {
			if bit&(1<<uint(i)) == 0 {
				node = node.left
			} else {
				node = node.right
			}
			if node.char != 0 {
				char := node.char

				sb.WriteRune(char)
				node = tree
			}
		}
	}

	return sb.String()
}
