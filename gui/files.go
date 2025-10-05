package gui

import (
	"fmt"
	"strings"

	"beckx.online/butils/fileutils"
	"beckx.online/yaatt/yaatt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
		ui.inheritTags()
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

	chkSelAll := widget.NewCheck("All", func(b bool) {
		ui.inheritTags()
		for _, tf := range ui.TheFiles {
			tf.Selected = b
		}
		files := fileutils.GetSelectedFilepathes(ui.TheFiles)
		tts := ui.yd.GetTextTags(files)
		ui.RefreshTagView(tts)
		ui.lstFiles.UnselectAll()
		ui.lstFiles.Refresh()
	})

	var btnInverseSelection *widget.Button
	btnInverseSelection = widget.NewButtonWithIcon("inverse",
		theme.MediaReplayIcon(), func() {
			ui.inheritTags()
			for _, tf := range ui.TheFiles {
				tf.Selected = !tf.Selected
			}
			files := fileutils.GetSelectedFilepathes(ui.TheFiles)
			tts := ui.yd.GetTextTags(files)
			ui.RefreshTagView(tts)
			ui.lstFiles.UnselectAll()
			ui.lstFiles.Refresh()
		})
	return container.NewVBox(lblFileheader,
		container.NewHBox(
			chkSelAll,
			btnInverseSelection))
}

func (ui *UI) inheritTags() {
	for _, fitem := range ui.frmTags.Items {
		yn := fitem.Text
		values := strings.Split(fitem.Widget.(*enterEntry).Text, yaatt.SEP_TAGVAL)
		fc := 0 // selected-filecount
		tc := 0 // tagvalue count: if filecount is greater then splited tags, last tagvalue will be inserted to all the rest files
		for _, file := range ui.TheFiles {
			if !file.Selected {
				continue
			}
			if len(values)-1 >= fc {
				tc = fc
			}
			md := ui.yd.MetaDatas[file.Path]
			fmt.Println("Length of MD->TextTags->yaatt-name", len(md.TextTags[yn]))
			if len(md.TextTags[yn]) == 0 {
				md.TextTags[yn] = []*yaatt.TextTag{
					{OrgName: ui.yd.Tagmap.YaatToId323[yn], Name: yn,
						Value: values[tc]},
				}
			} else {
				md.TextTags[yn][0].Value = values[tc]
			}
			fc++
		}
	}
}
