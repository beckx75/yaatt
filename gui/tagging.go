package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiTagging() *fyne.Container {
	tags := []string{
		"TAG-NAME",
		"artist", "album", "albartist",
		"cdnum", "tracknum", "year",
		"title", "genre", "comment",
	}

	frmItems := []*widget.FormItem{}

	frmTags := widget.NewForm()
	for _, tn := range tags {
		frmItems = append(frmItems, widget.NewFormItem(tn, widget.NewEntry()))
	}
	frmTags.Items = frmItems

	btnTest := widget.NewButton("Test...", func() {
		for _, frmItem := range frmItems {
			fmt.Println(frmItem.Widget.(*widget.Entry).Text)
		}
	})

	return container.NewBorder(nil, btnTest, nil, nil,
		frmTags,
	)
}

func (ui *UI) RefreshTags() {
	// save current tags

	// clear ui
	// insert new tags
	// refresh gui
}
