package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiFiles() *fyne.Container {
	selFiles := widget.NewSelect(ui.yd.Files, func(s string) {})

	tblFiles := widget.NewTable(
		func() (int, int) {
			return 100, 2
		},
		func() fyne.CanvasObject {
			return container.NewStack(widget.NewLabel("template11"), widget.NewLabel(""))
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*fyne.Container).Objects[0].(*widget.Label)
			e := o.(*fyne.Container).Objects[1].(*widget.Label)
			l.Show()
			e.Show()
			switch id.Col {
			case 0:
				e.Hide()
				l.SetText("hostname")
			case 1:
				l.Hide()
				e.SetText(fmt.Sprintf("Sepp %d", id.Row))
			}
		},
	)
	tblFiles.SetColumnWidth(0, 200)
	tblFiles.SetColumnWidth(1, 125)
	tblFiles.CreateHeader = func() fyne.CanvasObject {
		return widget.NewLabel("tmpl")
	}
	tblFiles.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {

		template.(*widget.Label).SetText("Header1")
	}
	tblFiles.ShowHeaderColumn = true
	tblFiles.OnSelected = func(id widget.TableCellID) {
		fmt.Println(id.Row, ":", id.Col)
	}

	// var lstFiles *widget.List
	// lstFiles = widget.NewList(
	// 	func() int { return 100 },
	// 	func() fyne.CanvasObject {
	// 		return container.NewStack(widget.NewLabel("template"), widget.NewEntry())
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// 		o.(*fyne.Container).Objects[0].(*widget.Label).SetText(fmt.Sprintf("%03d", i))
	// 		o.(*fyne.Container).Objects[1].(*widget.Entry).Text = fmt.Sprintf("%06x", i)
	// 	})
	// lstFiles.Refresh()

	return container.NewBorder(selFiles, nil, nil, nil,
		tblFiles,
	)
}
