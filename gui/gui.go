package gui

import (
	"bytes"
	"fmt"
	"huff/huffmann"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	window            fyne.Window
	footer            *widget.Label
	unencodedTextArea *widget.Entry
	encodedTextArea   *widget.Entry
	encodedBytes      []byte
	unencodedText     string
)

func Start() {
	myApp := app.New()
	window = myApp.NewWindow("Coder")
	window.Resize(fyne.NewSize(800, 600))

	unencodedTextArea = widget.NewMultiLineEntry()
	unencodedTextArea.SetPlaceHolder("Unencoded Text")
	unencodedTextArea.Wrapping = fyne.TextWrapWord

	encodedTextArea = widget.NewMultiLineEntry()
	encodedTextArea.SetPlaceHolder("Encoded Text")
	encodedTextArea.Wrapping = fyne.TextWrapWord
	encodedTextArea.Disable() // read-only

	encodeButton := widget.NewButton("Encode", func() {
		go encodedTextArea.SetText(encode(unencodedTextArea.Text))
	})

	decodeButton := widget.NewButton("Decode", func() {
		go unencodedTextArea.SetText(decode(encodedTextArea.Text))
	})

	leftTextStack := container.NewStack(unencodedTextArea)
	rightTextStack := container.NewStack(encodedTextArea)

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
	if len(text) == 0 {
		dialog.ShowError(fmt.Errorf("no text to encode"), window)
		return ""
	}

	encodedBytes = huffmann.Encode(text)
	var encodedText string

	for _, b := range encodedBytes {
		// display with leading zeros
		encodedText += fmt.Sprintf("%08b", b)
		encodedText += " "
	}

	setCompressionRatio(len(text), len(encodedBytes))

	return encodedText
}

func decode(text string) string {
	if len(text) == 0 {
		dialog.ShowError(fmt.Errorf("no text to decode"), window)
		return ""
	}

	var buffer bytes.Buffer

	for i := 0; i < len(text); i += 9 {
		b, _ := strconv.ParseUint(text[i:i+8], 2, 8)
		buffer.WriteByte(byte(b))
	}

	decodedText := huffmann.Decode(buffer.Bytes())

	setCompressionRatio(len(decodedText), len(buffer.Bytes()))

	encodedBytes = buffer.Bytes()
	unencodedText = text

	return decodedText
}
