package function

import "fyne.io/fyne/v2/widget"

func SetLoginPasswordData(loginPasswordData []string, loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry, loginEntry *widget.Entry, passwordEntry *widget.Entry) {
	loginPasswordNameEntry.SetText(loginPasswordData[0])
	loginPasswordDescriptionEntry.SetText(loginPasswordData[1])
	loginEntry.SetText(loginPasswordData[2])
	passwordEntry.SetText(loginPasswordData[3])
}

func SetTextData(textData []string, textNameEntry *widget.Entry, textDescriptionEntry *widget.Entry, textEntry *widget.Entry) {
	textNameEntry.SetText(textData[0])
	textDescriptionEntry.SetText(textData[1])
	textEntry.SetText(textData[2])
}
