package editor

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func DrawStr(msg string, x, y int, fg, bg termbox.Attribute) int {
	for _, char := range msg {
		x += DrawChar(char, x, y, fg, bg)
	}

	return x
}

func DrawChar(char rune, x, y int, fg, bg termbox.Attribute) int {
	termbox.SetCell(x, y, char, fg, bg)
	return runewidth.RuneWidth(char)
}
