package main

import (
	"altr/editor"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"os"
)

func main() {
	args := os.Args[1:]

	var name string
	switch len(args) {
	case 0:
		name = "untitled"
	case 1:
		name = args[0]
	default:
		fmt.Println("Expected 0 or 1 arguments.")
		os.Exit(0)
	}

	e, err := CreateEditor(name)
	if err != nil {
		panic(err)
	}
	defer e.Close()

	e.Draw()
	e.PollEvents()
}

func CreateEditor(name string) (*editor.Editor, error) {
	e := new(editor.Editor)
	e.Title = name
	e.Cursor = editor.Cursor{X: 0, Y: 1}
	e.Modified = false

	if name != "untitled" {
		file, err := os.Stat(name)
		if err == nil && !file.IsDir() {
			e.Content, err = readLines(name)
			if err != nil {
				e.Content = []string{}
			} else {
				e.IsNew = false
			}
		}
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	encoding.Register()
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	err = screen.Init()
	if err != nil {
		panic(err)
	}
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset))

	e.Screen = screen

	return e, nil
}
