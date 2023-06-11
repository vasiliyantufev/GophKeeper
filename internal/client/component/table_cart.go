package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SetDefaultColumnsWidthCart(table *widget.Table) {
	colWidths := []float32{150, 150, 150, 150, 50, 150, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func GetTableCart(dataTblCart [][]string) *widget.Table {
	tableDataCart := widget.NewTable(
		func() (int, int) {
			return len(dataTblCart), len(dataTblCart[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCart[i.Row][i.Col])
		})
	SetDefaultColumnsWidthCart(tableDataCart)
	return tableDataCart
}
