/*
 * Filename: /Users/arthur/go/src/gbtestapi/gbtestapi.go
 * Path: /Users/arthur/go/src/gbtestapi
 * Created Date: Wednesday, September 18th 2019, 3:13:21 pm
 * Author: Arthur Michajlenko
 *
 * Copyright (c) 2019 Infinit Loop SRL
 */

package main

import (
	"log"

	"fyne.io/fyne/dialog"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	log.Println("Start...")
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	win := app.NewWindow("Gelibert Test API")
	win.SetContent(widget.NewVBox(widget.NewLabel("Test API"),
		widget.NewEntry(),
		widget.NewButton("Quit", func() {
			dialog.NewConfirm("Quit", "Quit application...", func(ok bool) {
				if ok {
					app.Quit()
				}
			}, win).Show()
		}),
	))
	win.ShowAndRun()
}
