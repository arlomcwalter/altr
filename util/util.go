package util

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/mattn/go-runewidth"
	"math"
)

func NewScreen() tcell.Screen {
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

	return screen
}

func StrLen(msg string) int {
	var length int
	for _, char := range msg {
		length += runewidth.RuneWidth(char)
	}

	return length
}

func Min(min, val int) int {
	if val > min {
		return min
	}

	return val
}

func Max(max, val int) int {
	if val < max {
		return max
	}

	return val
}

func Clamp(min, max, val int) int {
	if val <= min {
		return min
	}

	if val >= max {
		return max
	}

	return val
}

func Abs(val int) int {
	if val > 0 {
		return val
	}

	return int(math.Abs(float64(val)))
}
