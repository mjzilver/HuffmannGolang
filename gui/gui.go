package gui

import (
	"bytes"
	"fmt"
	"huff/huffmann"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	byteSize      = 9 // 8 bits per byte + 1 space
	amountOfBytes = 1000
	maxLabelLen   = amountOfBytes * byteSize
)

var (
	window            fyne.Window
	footer            *widget.Label
	unencodedTextArea *widget.Entry
	encodedTextArea   *widget.Label
	encodedBytes      []byte
	unencodedText     string
)

func Start() {
	myApp := app.New()
	window = myApp.NewWindow("Huffman Coding in Go")
	window.Resize(fyne.NewSize(800, 600))

	unencodedTextArea = widget.NewMultiLineEntry()
	unencodedTextArea.SetPlaceHolder("Unencoded Text")
	unencodedTextArea.Wrapping = fyne.TextWrapWord

	encodedTextArea = widget.NewLabel("Encoded Text")
	encodedTextArea.Wrapping = fyne.TextWrapWord
	encodedTextArea.Resize(fyne.NewSize(400, 400))
	// Create a scrollable container for the encoded text area
	encodedScrollContainer := container.NewScroll(encodedTextArea)
	encodedScrollContainer.Resize(fyne.NewSize(400, 400))

	encodeButton := widget.NewButton("Encode", func() {
		if len(unencodedTextArea.Text) == 0 {
			dialog.ShowError(fmt.Errorf("no text to encode"), window)
			return
		}

		encodedTextArea.SetText("Loading...")
		go func() {
			encodedText := truncateLabel(encode(unencodedTextArea.Text))
			encodedTextArea.SetText(encodedText)
			window.Canvas().Refresh(encodedTextArea)
		}()
	})

	decodeButton := widget.NewButton("Decode", func() {
		if len(encodedBytes) == 0 {
			dialog.ShowError(fmt.Errorf("no encoded text to decode"), window)
			return
		}

		unencodedTextArea.SetText("Loading...")
		go func() {
			decodedText := decode(encodedTextArea.Text)
			unencodedTextArea.SetText(decodedText)
			window.Canvas().Refresh(unencodedTextArea)
		}()
	})

	leftTextStack := container.NewStack(unencodedTextArea)
	rightTextStack := container.NewStack(encodedScrollContainer)

	openTextFileButton := widget.NewButton("Open Text File", loadTextFromFile)
	saveBinaryFileButton := widget.NewButton("Save Encoded Text", saveEncodedTextToFile)
	openBinaryFileButton := widget.NewButton("Open Encoded Text File", loadBinaryFromFile)

	leftFooter := container.NewHBox(encodeButton, openTextFileButton)
	rightFooter := container.NewHBox(decodeButton, saveBinaryFileButton, openBinaryFileButton)

	leftContainer := container.NewBorder(nil, leftFooter, nil, nil, leftTextStack)
	rightContainer := container.NewBorder(nil, rightFooter, nil, nil, rightTextStack)

	footer = widget.NewLabel("Start by encoding some text.")

	textAreas := container.NewHSplit(leftContainer, rightContainer)
	content := container.NewBorder(nil, footer, nil, nil, textAreas)

	window.SetContent(content)
	window.ShowAndRun()
}

func encode(text string) string {
	encodedBytes = huffmann.Encode(text)
	var buffer bytes.Buffer
	for _, b := range encodedBytes {
		buffer.WriteString(fmt.Sprintf("%08b ", b))
	}

	encodedText := buffer.String()

	setCompressionRatio(len(text), len(encodedBytes))

	return encodedText
}

func decode(text string) string {
	decodedText := huffmann.Decode(encodedBytes)

	setCompressionRatio(len(decodedText), len(encodedBytes))
	unencodedText = text

	return decodedText
}

func truncateLabel(text string) string {
	if len(text) > maxLabelLen {
		return text[:maxLabelLen] + "..."
	}
	return text
}
