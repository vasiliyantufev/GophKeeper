package component

import "fyne.io/fyne/v2/widget"

func GetFormText(textName *widget.Entry, text *widget.Entry, textDescription *widget.Entry) *widget.Form {
	formText := widget.NewForm(
		widget.NewFormItem("Name", textName),
		widget.NewFormItem("Text", text),
		widget.NewFormItem("Description", textDescription),
	)
	return formText
}
