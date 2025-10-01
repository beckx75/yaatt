package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *UI) initUiTagging() *fyne.Container {
	// frmItems := []*widget.FormItem{}
	lblTags := widget.NewLabel("Tag's...")
	lblTags.Alignment = fyne.TextAlignCenter
	lblTags.TextStyle.Bold = true

	ui.frmTags = widget.NewForm()
	// for _, tn := range []string{} {
	// 	frmItems = append(frmItems, widget.NewFormItem(tn, widget.NewEntry()))
	// }
	// ui.frmTags.Items = frmItems

	btnTest := widget.NewButton("Test...", func() {
	})

	return container.NewBorder(lblTags, btnTest, nil, nil,
		ui.frmTags,
	)
}

func (ui *UI) RefreshTagView(tags [][]string) {
	// fmt.Println(tags)
	// save current tags
	// clear ui
	frmItems := []*widget.FormItem{}
	ui.frmTags.Items = frmItems
	ui.frmTags.Refresh()
	for _, tagsRow := range tags {
		entry := widget.NewEntry()
		entry.Text = tagsRow[1]
		frmItems = append(frmItems, widget.NewFormItem(tagsRow[0], entry))
	}
	ui.frmTags.Items = frmItems
	ui.frmTags.Refresh()
	// insert new tags
	// refresh gui
}
