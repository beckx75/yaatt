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

	fmt.Println(tags)

	tblTagging := widget.NewList(
		func() int {
			return len(tags)
		},
		func() fyne.CanvasObject {
			lblTagname := widget.NewLabel("tmpl")
			lblTagvalue := widget.NewLabel("tmpl")
			entValue := widget.NewEntry()

			return container.NewGridWithColumns(3,
				lblTagname, lblTagvalue, entValue,
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*fyne.Container).Objects[0].(*widget.Label).Text = tags[id]
		},
	)

	btnTest := widget.NewButton("test", func() {

	})

	return container.NewBorder(nil, btnTest, nil, nil,
		tblTagging,
	)
}
