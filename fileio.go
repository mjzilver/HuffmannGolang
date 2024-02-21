package main

import (
	"bufio"
	"encoding/binary"
	"log"
	"os"
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

func saveEncodedTextToFile(filename string, bytes []byte) int64 {
	file, err := os.Create(folder + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

func saveTextToFile(filename string, text string) {
	file, err := os.Create(folder + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}
}
