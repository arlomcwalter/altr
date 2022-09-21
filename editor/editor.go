package editor

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	Screen   tcell.Screen
	Title    string
	Content  []string
	Cursor   Cursor
	IsNew    bool
	Modified bool
}

var (
	// Colors
	accentColor = tcell.NewRGBColor(150, 76, 224)
	textColor   = tcell.ColorWhite
	invertColor = tcell.ColorSlateGray

	// Styles
	resetStyle     = tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	titleBarStyle  = tcell.StyleDefault.Foreground(textColor).Background(accentColor)
	statusBarStyle = tcell.StyleDefault.Foreground(textColor).Background(invertColor)
)

func (e *Editor) Close() {
	e.Screen.Fini()
}

func (e *Editor) PollEvents() {
	events := make(chan tcell.Event)

	go func() {
		for {
			if e.Screen == nil {
				break
			}
			events <- e.Screen.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-events:
			switch event := ev.(type) {
			case *tcell.EventKey:
				switch event.Key() {
				case tcell.KeyCtrlC:
					break loop
				case tcell.KeyLeft:
					e.Cursor.MoveX(-1, e)
					e.Draw()
				case tcell.KeyRight:
					e.Cursor.MoveX(1, e)
					e.Draw()
				case tcell.KeyUp:
					e.Cursor.MoveY(-1, e)
					e.Draw()
				case tcell.KeyDown:
					e.Cursor.MoveY(1, e)
					e.Draw()
				}
			case *tcell.EventResize:
				e.Draw()
			case *tcell.EventError:
				panic(event.Error())
			}
		}
	}
}

func (e *Editor) Draw() {
	e.Screen.ShowCursor(e.Cursor.X, e.Cursor.Y)
	e.Screen.Clear()
	e.Screen.Sync()

	width, height := e.Screen.Size()

	e.printTitleBar(width)
	e.printContent(width, height)
	e.printStatusBar(width, height)

	e.Screen.Show()
}

func (e *Editor) printTitleBar(width int) {
	currX := e.drawStr(" ALTR", 0, 0, titleBarStyle)

	titleLen := len(e.Title)
	halfWidth := (width - titleLen) / 2

	for currX < halfWidth {
		currX += e.drawChar(' ', currX, 0, titleBarStyle)
	}

	for currX < halfWidth+titleLen {
		currX += e.drawChar(rune(e.Title[currX-halfWidth]), currX, 0, titleBarStyle)
	}

	for currX < width {
		currX += e.drawChar(' ', currX, 0, titleBarStyle)
	}
}

func (e *Editor) printContent(width, height int) {
	lines := len(e.Content)

	for y := 1; y < height; y++ {
		relativeI := y - 1

		if lines <= relativeI {
			continue
		}

		line := e.Content[relativeI]
		if strLen(line) > width {
			line = trimStr(line, width)
			e.drawChar('…', width-1, y, resetStyle)
		}

		e.drawStr(line, 0, y, resetStyle)
	}
}

func (e *Editor) printStatusBar(width, height int) {
	message := fmt.Sprintf(" %d:%d (%d:%d)", e.Cursor.Y-1, e.Cursor.X, e.getHeight(), e.getWidth())

	currX := e.drawStr(message, 0, height-1, statusBarStyle)

	for currX < width {
		currX += e.drawChar(' ', currX, height-1, statusBarStyle)
	}
}

func (e *Editor) getWidth() int {
	var longest int

	for _, str := range e.Content {
		length := strLen(str)
		if length > longest {
			longest = length
		}
	}

	width, _ := e.Screen.Size()
	return max(width, longest)
}

func (e *Editor) getHeight() int {
	_, height := e.Screen.Size()
	return max(height-2, len(e.Content))
}
