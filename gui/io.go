package gui

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func loadTextFromFile() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		// cancel button was pressed
		if reader == nil {
			return
		}

		// show loading
		unencodedTextArea.SetText("Loading...")

		content, err := io.ReadAll(reader)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		unencodedTextArea.SetText(string(content))
	}, window)
}

func saveEncodedTextToFile() {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		if !strings.HasSuffix(writer.URI().Path(), ".bin") {
			dialog.ShowError(fmt.Errorf("file must have .bin extension"), window)
		}

		err = binary.Write(writer, binary.LittleEndian, encodedBytes)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		err = writer.Close()
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
	}, window)
}

func loadBinaryFromFile() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		// cancel button was pressed
		if reader == nil {
			return
		}
		encodedTextArea.SetText("Loading...")

		if reader.URI().Extension() != ".bin" {
			dialog.ShowError(fmt.Errorf("file is not binary"), window)
			return
		}

		content, err := io.ReadAll(reader)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		encodedBytes = content
		var encodedText string

		for _, b := range encodedBytes {
			// display with leading zeros
			encodedText += fmt.Sprintf("%08b", b)
			encodedText += " "
		}

		encodedText = truncateLabel(encodedText)
		encodedTextArea.SetText(encodedText)
	}, window)
}

func setCompressionRatio(originalFileSize, encodedFileSize int) {
	compressionRatio := float64(originalFileSize) / float64(encodedFileSize)
	footer.SetText(fmt.Sprintf("Compression ratio: %.2f", compressionRatio))
}
