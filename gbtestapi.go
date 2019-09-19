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

	"github.com/rivo/tview"
)

func main() {
	log.Println("Start...")
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	// app:=tview.NewApplication()
	if err:= tview.NewApplication().SetRoot(box,true).Run(); err != nil {
		panic(err)
	}
}
