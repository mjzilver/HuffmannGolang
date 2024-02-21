package main

import (
	"time"
)

func main() {
	originalText, originalFileSize := loadTextFromFile("input.txt")

	startTime := time.Now()

	// add EOF character to the text
	text := originalText + string(EOF)

	freq := countFrequency(text)
	tree := buildHuffmanTree(freq)
	codes := generateCodes(tree)

	encodedText := encode(text, codes)
	decodedText := decode(encodedText, tree)

	saveTextToFile("decoded.txt", decodedText)

	elapsedTime := time.Since(startTime)

	// output debug information if the original text and the decoded text are not the same
	if originalText != decodedText {
		println("Original text:", originalText)
		println("Encoded text:", encodedText)
		println("Decoded text:", decodedText)
		println("The original text and the decoded text are not the same")
		println("Original text length:", len(text))
		println("Decoded text length:", len(decodedText))

		// print last 10 characters of the original and the decoded text
		println("Last 10 characters of the original text:", text[len(text)-10:])
		println("Last 10 characters of the decoded text:", decodedText[len(decodedText)-10:])

		// print the part of the text where the original and the decoded text differ
		for i := 0; i < len(text); i++ {
			if text[i] != decodedText[i] {
				println("Part of original text", text[i-10:i])
				println("Part of decoded text:", decodedText[i-10:i])
				break
			}
		}
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
