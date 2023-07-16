package function

import "fyne.io/fyne/v2/widget"

func ClearText(textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry) {
	textNameEntry.SetText("")
	textEntry.SetText("")
	textDescriptionEntry.SetText("")
}

func ClearCard(cardNameEntry *widget.Entry, cardDescriptionEntry *widget.Entry, paymentSystemEntry *widget.Entry,
	numberEntry *widget.Entry, holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry) {
	cardNameEntry.SetText("")
	cardDescriptionEntry.SetText("")
	paymentSystemEntry.SetText("")
	numberEntry.SetText("")
	holderEntry.SetText("")
	endDateEntry.SetText("")
	cvcEntry.SetText("")
}

func ClearLoginPassword(loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry,
	loginEntry *widget.Entry, passwordEntry *widget.Entry) {
	loginPasswordNameEntry.SetText("")
	loginPasswordDescriptionEntry.SetText("")
	loginEntry.SetText("")
	passwordEntry.SetText("")
}
