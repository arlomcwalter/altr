package editor

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

type Editor struct {
	title   string
	content []string
	offX    int
	offY    int
	Cursor  Cursor
}

var editorName = "ALTR"

var accentColor = termbox.RGBToAttribute(150, 76, 224)
var textColor = termbox.RGBToAttribute(255, 255, 255)

func CreateEditor() *Editor {
	name := "untitled"
	var content []string

	if len(os.Args[1:]) > 0 {
		name = os.Args[1]

		file, err := os.Stat(name)
		if err == nil && !file.IsDir() {
			name = file.Name()
			content, err = readLines(name)
			if err != nil {
				content = []string{}
			}
		}
	}

	editor := new(Editor)
	editor.title = name
	editor.content = content
	editor.Cursor.Set(0, 1, editor)

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
			switch event.Type {
			case termbox.EventKey:
				if event.Ch != 0 {

				} else if event.Key != 0 {
					switch event.Key {
					case termbox.KeyCtrlC:
						break loop
					case termbox.KeyArrowLeft:
						e.Cursor.MoveX(-1, e)
						e.Draw()
					case termbox.KeyArrowRight:
						e.Cursor.MoveX(1, e)
						e.Draw()
					case termbox.KeyArrowUp:
						e.Cursor.MoveY(-1, e)
						e.Draw()
					case termbox.KeyArrowDown:
						e.Cursor.MoveY(1, e)
						e.Draw()
					}
				}
			case termbox.EventResize:
				e.Draw()
			case termbox.EventError:
				panic(event.Err)
			}
		}
	}
}

func (e *Editor) Draw() {
	termbox.SetCursor(e.Cursor.x, e.Cursor.y)

	err := termbox.Sync()
	if err != nil {
		panic(err)
	}

	width, height := termbox.Size()

	err = termbox.Clear(textColor, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	e.printTitleBar(width)
	e.printContent(width, height)
	e.printStatusBar(width, height)

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (e *Editor) printTitleBar(width int) {
	currX := drawStr(" "+editorName, 0, 0, textColor, accentColor)

	titleLen := len(e.title)
	halfWidth := (width - titleLen) / 2

	for currX < halfWidth {
		currX += drawChar(' ', currX, 0, textColor, accentColor)
	}

	for currX < halfWidth+titleLen {
		currX += drawChar(rune(e.title[currX-halfWidth]), currX, 0, textColor, accentColor)
	}

	for currX < width {
		currX += drawChar(' ', currX, 0, textColor, accentColor)
	}
}

func (e *Editor) printContent(width, height int) {
	lines := len(e.content)

	for y := 1; y < height; y++ {
		relativeI := y - 1

		if lines <= relativeI {
			continue
		}

		line := e.content[relativeI]
		if strLen(line) > width {
			line = trimStr(line, width)
			drawChar('…', width-1, y, textColor, termbox.ColorDefault)
		}

		drawStr(line, 0, y, textColor, termbox.ColorDefault)
	}
}

func (e *Editor) printStatusBar(width, height int) {
	message := fmt.Sprintf(" %d:%d (%d:%d)", e.Cursor.y-1, e.Cursor.x, e.docHeight(), e.docWidth())

	currX := drawStr(message, 0, height-1, termbox.ColorBlack, textColor)

	for currX < width {
		currX += drawChar(' ', currX, height-1, termbox.ColorBlack, textColor)
	}
}
