package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiFiles() *fyne.Container {
	var lstFiles *widget.List
	lstFiles = widget.NewList(
		func() int { return 7 },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).Text = fmt.Sprintf("%d", i)
		})
	lstFiles.Refresh()

	return container.NewBorder(nil, nil, nil, nil,
		lstFiles,
	)
}
