package main

import (
	"altr/editor"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.OutputRGB)

	e := editor.CreateEditor()
	e.Draw()
	e.PollEvents()
}
