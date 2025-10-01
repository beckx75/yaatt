package gui

import (
	"beckx.online/butils/fileutils"
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
		// save current tagitems-data
		// TODO SAVE DATA..

		ui.TheFiles[id].Selected = !ui.TheFiles[id].Selected
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

	var btnSelectAllFiles *widget.Button
	btnSelectAllFiles = widget.NewButton("select all", func() {
		if !ui.selectAllButtonPressed {
			for i, tf := range ui.TheFiles {
				tf.Selected = true
				ui.lstFiles.Select(i)
				ui.lstFiles.OnSelected(i)
			}
			ui.lstFiles.UnselectAll()
			ui.lstFiles.Refresh()
			ui.selectAllButtonPressed = true
			btnSelectAllFiles.Text = "de-select all"
			btnSelectAllFiles.Refresh()
		} else {
			// deselect everything
			for i, tf := range ui.TheFiles {
				tf.Selected = false
				ui.lstFiles.Select(i)
				ui.lstFiles.OnSelected(i)
			}
			ui.lstFiles.UnselectAll()
			ui.lstFiles.Refresh()
			ui.selectAllButtonPressed = false
			btnSelectAllFiles.Text = "select all"
			btnSelectAllFiles.Refresh()
		}
		// NEXT update lstFiles with all files selected
	})

	var btnInverseSelection *widget.Button
	btnInverseSelection = widget.NewButton("inverse selection", func() {
		for i, tf := range ui.TheFiles {
			tf.Selected = !tf.Selected
			ui.lstFiles.Select(i)
			ui.lstFiles.OnSelected(i)
		}
		ui.lstFiles.UnselectAll()
		ui.lstFiles.Refresh()
	})
	return container.NewVBox(lblFileheader,
		container.NewHBox(btnSelectAllFiles, btnInverseSelection))
}
