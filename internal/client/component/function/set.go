package function

import "fyne.io/fyne/v2/widget"

func SetLoginPasswordData(loginPasswordData []string, loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry, loginEntry *widget.Entry, passwordEntry *widget.Entry) {
	loginPasswordNameEntry.SetText(loginPasswordData[0])
	loginPasswordDescriptionEntry.SetText(loginPasswordData[1])
	loginEntry.SetText(loginPasswordData[2])
	passwordEntry.SetText(loginPasswordData[3])
}
