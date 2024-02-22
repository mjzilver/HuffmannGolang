package main

import (
	"time"
)

func main() {
	text, originalFileSize := loadTextFromFile("input.txt")

	startTime := time.Now()

	freq := countFrequency(text)
	tree := buildHuffmanTree(freq)
	codes := generateCodes(tree)

	encodedText := encode(text, codes)
	decodedText := decode(encodedText, tree)

	saveTextToFile("decoded.txt", decodedText)

	elapsedTime := time.Since(startTime)

	// output debug information if the original text and the decoded text are not the same
	if text != decodedText {
		println("Original text:", text)
		println("Encoded text:", encodedText)
		println("Decoded text:", decodedText)
		println("The original text and the decoded text are not the same")
	}

	// encode the tree
	encodedTree := encodeTree(tree)

	fileBytes := append(encodedText, encodedTree...)

	// save the encoded text to a file
	encodedFileSize := saveEncodedTextToFile("encoded.bin", fileBytes)

	println("Original text size:", originalFileSize)
	println("Encoded text size:", encodedFileSize)

	// calculate compression ratio
	compressionRatio := float64(originalFileSize) / float64(encodedFileSize)
	println("Compression ratio:", compressionRatio)
	if elapsedTime.Milliseconds() == 0 {
		println("Elapsed microseconds:", elapsedTime.Microseconds())
	} else {
		println("Elapsed milliseconds:", elapsedTime.Milliseconds())
	}

	decodedTree := decodeTree(encodedTree)

	println("Original tree", tree.String())
	println("Decoded tree ", decodedTree.String())
}
