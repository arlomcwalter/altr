package editor

import (
	"github.com/nsf/termbox-go"
	"os"
)

type Editor struct {
	title   string
	content []string
}

var editorName = "ALTR"

func CreateEditor() *Editor {
	editor := new(Editor)

	fileName := "untitled"
	if len(os.Args[1:]) > 0 {
		fileName = os.Args[1]
	}
	editor.title = fileName

	return editor
}

func (e *Editor) PollEvents() {
	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case event := <-events:
			if event.Type == termbox.EventKey {
				if event.Key == termbox.KeyCtrlC {
					break loop
				}
			}
			if event.Type == termbox.EventResize {
				e.Draw()
			}
		}
	}
}

func (e *Editor) Draw() {
	width, _ := termbox.Size()

	err := termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
	if err != nil {
		return
	}

	e.drawStatusBar(width)

	err = termbox.Flush()
	if err != nil {
		return
	}
}

func (e *Editor) drawStatusBar(width int) {
	var currX = DrawStr(" "+editorName, 0, 0, termbox.ColorWhite, termbox.ColorBlue)

	titleLen := len(e.title)
	halfWidth := (width - titleLen) / 2

	for currX < halfWidth {
		currX += DrawChar(' ', currX, 0, termbox.ColorWhite, termbox.ColorBlue)
	}

	for currX < halfWidth+titleLen {
		currX += DrawChar(rune(e.title[currX-halfWidth]), currX, 0, termbox.ColorWhite, termbox.ColorBlue)
	}

	for currX < width {
		currX += DrawChar(' ', currX, 0, termbox.ColorWhite, termbox.ColorBlue)
	}
}
