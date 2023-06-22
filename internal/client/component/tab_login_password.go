package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetTabLoginPassword(tblLoginPassword *widget.Table, top *widget.Button, loginPassword *widget.Button) *container.TabItem {
	containerTblLoginPassword := layout.NewBorderLayout(top, loginPassword, nil, nil)
	boxLoginPassword := fyne.NewContainerWithLayout(containerTblLoginPassword, top, tblLoginPassword, loginPassword)
	return container.NewTabItem("Логин пароль", boxLoginPassword)
}
