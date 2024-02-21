package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	folder = "files/"
)

func loadTextFromFile(filename string) (string, int64) {
	file, err := os.Open(folder + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text, fileSize
}

func saveEncodedTextToFile(filename string, encodedText string) int64 {
	file, err := os.Create(folder + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bytes []byte
	for i := 0; i < len(encodedText); i += 8 {
		end := i + 8
		if end > len(encodedText) {
			end = len(encodedText)
		}
		binaryChunk := encodedText[i:end]
		binaryValue, err := strconv.ParseInt(binaryChunk, 2, 16)
		if err != nil {
			fmt.Println("Error parsing binary string:", err)
			return 0
		}
		bytes = append(bytes, byte(binaryValue))
	}

	err = binary.Write(file, binary.LittleEndian, bytes)
	if err != nil {
		log.Fatal(err)
	}

	// calculate the file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()

	return fileSize
}
