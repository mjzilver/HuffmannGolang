package huffmann

const (
	EOF    = rune(0x00)
	marker = byte(0xFF)
)

func countFrequency(text string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}
	return freq
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
