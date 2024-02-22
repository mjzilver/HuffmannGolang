package main

import (
	"huff/huffmann"
	"time"
)

func main() {
	text, originalFileSize := loadTextFromFile("input.txt")

	startTime := time.Now()

	encodedText := huffmann.Encode(text)
	encodedFileSize := saveEncodedTextToFile("encoded.bin", encodedText)

	decodedText := huffmann.Decode(encodedText)

	saveTextToFile("decoded.txt", decodedText)

	elapsedTime := time.Since(startTime)

	// output debug information if the original text and the decoded text are not the same
	if text != decodedText {
		println("Original text:", text)
		println("Encoded text:", encodedText)
		println("Decoded text:", decodedText)
		println("The original text and the decoded text are not the same")

		println("Original text length:", len(text))
		println("Decoded text length:", len(decodedText))
	}

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
