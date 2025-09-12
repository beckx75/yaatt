package gui

import (
	"fmt"

	"beckx.online/yaatt/yaatt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app fyne.App
	win fyne.Window

	yd *yaatt.YaattData
}

func InitGui(args []string) {
	var err error
	ui := &UI{}

	ui.app = app.New()
	ui.win = ui.app.NewWindow("this is yaatt...")

	if len(args) > 0 {
		ui.yd, err = yaatt.NewYaattData(args, ".")
		if err != nil {
			dialog.ShowError(err, ui.win)
		} else {
			dialog.ShowInformation("args success",
				fmt.Sprintf("read in %d audiofiles", len(ui.yd.MetaDatas)),
				ui.win,
			)
		}
	}

	// var err error

	// cntFiles := ui.initUiFiles()
	// cntFrames := ui.makeTagcontent()

	// toolbar := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.HomeIcon(), func() { fmt.Println("feeling like home...") }),
	// 	widget.NewToolbarSpacer(),
	// 	widget.NewToolbarAction(theme.LogoutIcon(), func() { ui.app.Quit() }),
	// )

	cntFiles := ui.initUiFiles()

	btnQuit := widget.NewButtonWithIcon("i'm done...", theme.HomeIcon(), func() {
		ui.app.Quit()
	})

	mainbox := container.NewBorder(
		cntFiles, btnQuit, nil, nil,
		widget.NewLabel("coming soon..."),
	)

	// ui.win.SetOnDropped(
	// 	func(p fyne.Position, uris []fyne.URI) {
	// 		for _, uri := range uris {
	// 			fmt.Println(uri)
	// 			fv := fileutils.NewTheFile(uri.Path())
	// 			ui.thefiles = append(ui.thefiles, fv)
	// 		}
	// 	})

	ui.win.SetContent(mainbox)
	ui.win.Resize(fyne.NewSize(1024, 768))
	ui.win.CenterOnScreen()
	ui.win.ShowAndRun()
}
