package window

import (
	"altr/util"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// Data structure

type Window struct {
	title  string
	screen tcell.Screen
	editor *Editor
	cursor *Cursor
	events chan tcell.Event
}

func Create(name string) *Window {
	window := new(Window)
	window.title = name
	window.screen = util.NewScreen()
	window.editor = createEditor(util.Parse(name))
	window.cursor = &Cursor{col: 0, line: 0}
	window.events = make(chan tcell.Event)
	go func() {
		for {
			window.events <- window.screen.PollEvent()
		}
	}()

	return window
}

// Main interface

func (w *Window) Init() {
	w.screen.Sync()
	w.update()
	w.pollEvents()
}

func (w *Window) Shutdown() {
	w.screen.Fini()
}

// Functionality

func (w *Window) pollEvents() {
loop:
	for {
		select {
		case ev := <-w.events:
			switch event := ev.(type) {
			case *tcell.EventKey:
				switch event.Key() {
				case tcell.KeyCtrlC:
					break loop
				case tcell.KeyLeft:
					w.cursor.moveX(-1)
					w.update()
				case tcell.KeyRight:
					w.cursor.moveX(1)
					w.update()
				case tcell.KeyUp:
					w.cursor.moveY(-1)
					w.update()
				case tcell.KeyDown:
					w.cursor.moveY(1)
					w.update()
				}
			case *tcell.EventResize:
				w.update()
			case *tcell.EventError:
				panic(event.Error())
			}
		}
	}
}

func (w *Window) update() {
	w.screen.Clear()

	cols, lines := w.screen.Size()

	w.editor.update(lines)
	w.cursor.clamp(cols, lines, w.editor)
	w.screen.ShowCursor(w.cursor.pos(w.editor))

	//draw
	w.printTitleBar(cols)
	w.editor.draw(w, w.cursor.line+w.editor.scrollY)
	w.printStatusBar(cols, lines)

	w.screen.Show()
}

// Window elements

func (w *Window) printTitleBar(cols int) {
	currX := w.drawStr(" ALTR", 0, 0, titleBarStyle)

	titleLen := len(w.title)
	halfWidth := (cols - titleLen) / 2

	for currX < halfWidth {
		currX += w.drawChar(' ', currX, 0, titleBarStyle)
	}

	for currX < halfWidth+titleLen {
		currX += w.drawChar(rune(w.title[currX-halfWidth]), currX, 0, titleBarStyle)
	}

	for currX < cols {
		currX += w.drawChar(' ', currX, 0, titleBarStyle)
	}
}

func (w *Window) printStatusBar(cols, lines int) {
	message := fmt.Sprintf(" %d:%d (%d lines)", w.cursor.line+w.editor.scrollY+1, w.cursor.col+w.editor.scrollX+1, w.editor.lineCount())

	currX := w.drawStr(message, 0, lines-1, statusBarStyle)

	for currX < cols {
		currX += w.drawChar(' ', currX, lines-1, statusBarStyle)
	}
}

// Rendering

func (w *Window) drawStr(msg string, x, y int, style tcell.Style) int {
	for _, char := range msg {
		x += w.drawChar(char, x, y, style)
	}

	return x
}

func (w *Window) drawChar(char rune, x, y int, style tcell.Style) int {
	w.screen.SetContent(x, y, char, nil, style)
	return runewidth.RuneWidth(char)
}
