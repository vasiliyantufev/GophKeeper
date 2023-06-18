package form

import "fyne.io/fyne/v2/widget"

func ClearText(textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry) {
	textNameEntry.SetText("")
	textEntry.SetText("")
	textDescriptionEntry.SetText("")
}

func ClearCart(cartNameEntry *widget.Entry, paymentSystemEntry *widget.Entry, numberEntry *widget.Entry,
	holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry) {
	cartNameEntry.SetText("")
	paymentSystemEntry.SetText("")
	numberEntry.SetText("")
	holderEntry.SetText("")
	endDateEntry.SetText("")
	cvcEntry.SetText("")
}
