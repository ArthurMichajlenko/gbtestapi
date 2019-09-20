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
	"strconv"

	"github.com/rivo/tview"
)

func main() {
	log.Println("Start...")
	app := tview.NewApplication()
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Test API")
	form.AddInputField("ID", "", 10, tview.InputFieldInteger, nil).
		AddInputField("Name", "", 20, nil, nil)
	form.AddButton("Ok", func() {
		app.Stop()
		id, _ := strconv.Atoi(form.GetFormItemByLabel("ID").(*tview.InputField).GetText())
		log.Println(form.GetFormItemByLabel("ID").(*tview.InputField).GetText(), id)
		log.Println(form.GetFormItemByLabel("Name").(*tview.InputField).GetText())
	})
	if err := app.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}
}
