package gui

import (
	"fmt"
	"path/filepath"

	"beckx.online/yaatt/yaatt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
)

type TheFile struct {
	Path     string
	Name     string
	Dir      string
	Selected bool
}

type UI struct {
	app fyne.App
	win fyne.Window

	lstFiles *widget.List

	yd       *yaatt.YaattData
	TheFiles []TheFile

	bindFileHeader binding.String
}

func InitGui(args []string) {
	var err error
	ui := &UI{
		TheFiles:       []TheFile{},
		bindFileHeader: binding.NewString(),
	}

	ui.app = app.New()
	ui.win = ui.app.NewWindow("this is yaatt...")

	if len(args) > 0 {
		ui.yd, err = yaatt.NewYaattData(args, ".")
		if err != nil {
			// dialog.ShowError(err, ui.win)
			log.Warn().Msgf("%v", err)
		}
		ui.TheFiles = MakeTheFileList(ui.yd.Files)
		ui.bindFileHeader.Set(fmt.Sprintf("Files (%d)", len(ui.TheFiles)))

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
	cntTagging := ui.initUiTagging()

	lblYaatt := widget.NewLabel("Enjoy and have fun...")
	lblYaatt.Alignment = fyne.TextAlignCenter
	lblYaatt.TextStyle.Bold = true
	cntContent := container.NewBorder(lblYaatt, nil, nil, nil,
		container.NewGridWithColumns(2,
			cntFiles, cntTagging,
		),
	)

	btnQuit := widget.NewButtonWithIcon("i'm done...", theme.HomeIcon(), func() {
		ui.app.Quit()
	})

	mainbox := container.NewBorder(
		nil, btnQuit, nil, nil,
		cntContent,
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

func MakeTheFile(fp string) TheFile {
	return TheFile{
		Path:     fp,
		Name:     filepath.Base(fp),
		Dir:      filepath.Dir(fp),
		Selected: false,
	}
}

func MakeTheFileList(fps []string) []TheFile {
	lst := []TheFile{}
	for _, fp := range fps {
		lst = append(lst, MakeTheFile(fp))
	}
	return lst
}

func MakeFilenameList(tfs []TheFile) []string {
	lst := []string{}
	for _, tf := range tfs {
		lst = append(lst, tf.Name)
	}
	return lst
}
