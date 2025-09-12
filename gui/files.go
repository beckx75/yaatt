package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiFiles() *fyne.Container {
	selFiles := widget.NewSelect(ui.yd.Files, func(s string) {})

	return container.NewBorder(selFiles, nil, nil, nil)
}
