package window

import (
	"altr/util"
	"fmt"
	"strconv"
)

// Data structure

type Editor struct {
	visibleLines     int
	padT, padB       int
	padL, padR       int
	scrollX, scrollY int
	content          []string
}

func createEditor(content []string) *Editor {
	editor := new(Editor)
	editor.content = content
	editor.padT = 1
	editor.padB = 1
	return editor
}

// Dimensions

func (e *Editor) update(visibleLines int) {
	e.padL = 1 + len(strconv.Itoa(e.lineCount())) + 1
	e.visibleLines = visibleLines - e.vPadding()
}

func (e *Editor) calcWidthAt(line int) int {
	actual := line + e.scrollY
	if actual >= e.lineCount() {
		return 0
	}

	content := e.content[actual]
	if len(content) > e.scrollX {
		content = content[e.scrollX:]
	}

	return util.StrWidth(content)
}

func (e *Editor) lineCount() int {
	return len(e.content)
}

func (e *Editor) vPadding() int {
	return e.padT + e.padB
}

// Rendering

func (e *Editor) draw(window *Window, active int) {
	for line := 1; line <= e.visibleLines; line++ {
		fileLine := line + e.scrollY
		if fileLine > e.lineCount() {
			break
		}

		lineRef := fileLine - 1

		style := lineNumStyle
		if lineRef == active {
			style = activeLineStyle
		}

		absY := line - 1 + e.padT
		window.drawStr(fmt.Sprintf(" %d ", fileLine), 0, absY, style)

		content := e.content[lineRef]
		if len(content) > e.scrollX {
			window.drawStr(content[e.scrollX:], e.padL, absY, resetStyle)
		}
	}
}
