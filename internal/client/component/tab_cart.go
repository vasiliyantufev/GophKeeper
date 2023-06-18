package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetTabCarts(tblCart *widget.Table, top *widget.Button, cart *widget.Button) *container.TabItem {
	containerTblCart := layout.NewBorderLayout(top, cart, nil, nil)
	boxCart := fyne.NewContainerWithLayout(containerTblCart, top, tblCart, cart)
	return container.NewTabItem("Банковские карты", boxCart)
}
