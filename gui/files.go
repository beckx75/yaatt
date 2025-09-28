package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiFiles() *fyne.Container {
	ui.lstFiles = widget.NewList(
		func() int {
			return len(ui.TheFiles)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			lbl := o.(*widget.Label)
			lbl.TextStyle.Bold = ui.TheFiles[i].Selected
			lbl.SetText(ui.TheFiles[i].Name)
		})
	ui.lstFiles.OnSelected = func(id widget.ListItemID) {
		ui.TheFiles[id].Selected = !ui.TheFiles[id].Selected
		ui.lstFiles.UnselectAll()
		ui.lstFiles.Refresh()
		// save current tagitems-data
		// clear tagitems-view
		// get new selected file tagitems
		// fill new selected tagitems
		// show ne selected tagitems
	}
	cntFilesHeader := ui.makeFilesHeader()
	return container.NewBorder(cntFilesHeader, nil, nil, nil,
		ui.lstFiles,
	)
}

func (ui *UI) makeFilesHeader() *fyne.Container {
	lblFileheader := widget.NewLabelWithData(ui.bindFileHeader)
	lblFileheader.TextStyle.Bold = true
	lblFileheader.Alignment = fyne.TextAlignCenter

	btnSelectAllFiles := widget.NewButton("select all", func() {
		for _, tf := range ui.TheFiles {
			tf.Selected = !tf.Selected
		}
		// NEXT update lstFiles with all files selected
	})

	return container.NewVBox(lblFileheader, btnSelectAllFiles)
}
