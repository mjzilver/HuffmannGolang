package huffmann

import "bytes"

func Encode(text string) []byte {
	freq := countFrequency(text)
	tree := buildHuffmanTree(freq)
	codes := generateCodes(tree)

	var buffer bytes.Buffer
	var bits byte
	var bitCount int

	// start binary with the tree
	treeBytes := encodeTree(tree)
	buffer.Write(treeBytes)
	// end tree with one FF byte
	buffer.WriteByte(marker)

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

	// default padding is 0
	var padding uint8 = 0

	if bitCount > 0 {
		// Pad the remaining bits to form a complete byte
		bits = bits << (8 - bitCount)
		padding = uint8(8 - bitCount)
		buffer.WriteByte(bits)
	}

	// Add padding byte to the end of the file
	buffer.WriteByte(padding)

	return buffer.Bytes()
}
