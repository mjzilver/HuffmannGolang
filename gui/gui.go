package gui

import (
	"bytes"
	"fmt"
	"huff/huffmann"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Start() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Coder")
	myWindow.Resize(fyne.NewSize(800, 600))

	unencodedTextArea := widget.NewMultiLineEntry()
	unencodedTextArea.SetPlaceHolder("Unencoded Text")
	unencodedTextArea.Wrapping = fyne.TextWrapWord

	encodedTextArea := widget.NewMultiLineEntry()
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

	leftContainer := container.NewBorder(nil, encodeButton, nil, nil, leftTextStack)
	rightContainer := container.NewBorder(nil, decodeButton, nil, nil, rightTextStack)

	content := container.NewHSplit(leftContainer, rightContainer)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func encode(text string) string {
	bytes := huffmann.Encode(text)
	var encodedText string

	for _, b := range bytes {
		// display with leading zeros
		encodedText += fmt.Sprintf("%08b", b)
		encodedText += " "
	}

	return encodedText
}

func decode(text string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(text); i += 9 {
		b, _ := strconv.ParseUint(text[i:i+8], 2, 8)
		buffer.WriteByte(byte(b))
	}

	decodedText := huffmann.Decode(buffer.Bytes())

	return decodedText
}
