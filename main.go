package main

import (
	"log"
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

	elapsedTime := time.Since(startTime)

	if text != decodedText {
		log.Fatal("The original text and the decoded text are not the same")
	}

	// save the encoded text to a file
	encodedFileSize := saveEncodedTextToFile("encoded.bin", encodedText)

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
}
