package function

import "fyne.io/fyne/v2/widget"

func HideLabelsTab(labelAlertLoginPassword *widget.Label, labelAlertText *widget.Label, labelAlertCard *widget.Label) {
	labelAlertLoginPassword.Hide()
	labelAlertText.Hide()
	labelAlertCard.Hide()
}
