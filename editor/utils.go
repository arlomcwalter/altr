package editor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// Errors

type InvalidFileErr struct{}

func (m *InvalidFileErr) Error() string {
	return "Invalid file input."
}

// Drawing to screen

func (e *Editor) drawStr(msg string, x, y int, style tcell.Style) int {
	for _, char := range msg {
		x += e.drawChar(char, x, y, style)
	}

	return x
}

func (e *Editor) drawChar(char rune, x, y int, style tcell.Style) int {
	e.Screen.SetContent(x, y, char, nil, style)
	return runewidth.RuneWidth(char)
}

// Screen clipping

func strLen(msg string) int {
	var length int
	for _, char := range msg {
		length += runewidth.RuneWidth(char)
	}

	return length
}

func trimStr(msg string, width int) string {
	if strLen(msg) <= width {
		return msg
	}

	var length int
	for i, char := range msg {
		newChar := runewidth.RuneWidth(char)
		if length+newChar > width {
			return msg[0 : i-1]
		}

		length += newChar
	}

	return msg
}

// Math

func max(max, val int) int {
	if val > max {
		return max
	}

	return val
}

func clamp(min, max, val int) int {
	if val <= min {
		return min
	}

	if val >= max {
		return max
	}

	return val
}
