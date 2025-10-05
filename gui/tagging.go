package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type enterEntry struct {
	widget.Entry
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

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

	btnTest := widget.NewButtonWithIcon("Save",
		theme.DocumentSaveIcon(), func() {
			ui.inheritTags()
			err := ui.yd.WriteMetadata()
			if err != nil {
				dialog.ShowError(err, ui.win)
			}
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
		entry := newEnterEntry()
		entry.Text = tagsRow[1]
		frmItems = append(frmItems, widget.NewFormItem(tagsRow[0], entry))
	}
	ui.frmTags.Items = frmItems
	ui.frmTags.Refresh()
	// insert new tags
	// refresh gui
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn:
		fmt.Println(e.BaseWidget)
		fmt.Println(e.Position())
		e.FocusLost()
		fmt.Println("ja druggd der sauhund enta ha")
	default:
		e.Entry.KeyDown(key)
		// fmt.Printf("Key %v pressed\n", key.Name)
	}
}
