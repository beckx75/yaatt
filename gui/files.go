package gui

import (
	"beckx.online/butils/fileutils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiFiles() *fyne.Container {
	// ui.lstFiles = widget.NewList(
	// 	func() int {
	// 		return len(ui.TheFiles)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewLabel("template")
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// 		lbl := o.(*widget.Label)
	// 		lbl.TextStyle.Bold = ui.TheFiles[i].Selected
	// 		lbl.SetText(ui.TheFiles[i].Name)
	// 	})
	ui.lstFiles = widget.NewListWithData(
		ui.bindFiles,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			lbl := o.(*widget.Label)
			lbl.TextStyle.Bold = ui.TheFiles[i.(int)].Selected
			lbl.Bind(i.(binding.String))
		},
	)
	ui.lstFiles.OnSelected = func(id widget.ListItemID) {
		// save current tagitems-data
		// TODO SAVE DATA..

		ui.TheFiles[id].Selected = !ui.TheFiles[id].Selected
		ui.lstFiles.
			ui.lstFiles.UnselectAll()
		ui.lstFiles.Refresh()
		// clear tagitems-view
		// get new selected file tagitems
		files := fileutils.GetSelectedFilepathes(ui.TheFiles)
		tts := ui.yd.GetTextTags(files)
		// fill new selected tagitems
		ui.RefreshTagView(tts)
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
		for i, tf := range ui.TheFiles {
			tf.Selected = true
			ui.lstFiles.Select(i)
		}
		ui.lstFiles.UnselectAll()
		ui.lstFiles.Refresh()
		// NEXT update lstFiles with all files selected
	})

	return container.NewVBox(lblFileheader, btnSelectAllFiles)
}
