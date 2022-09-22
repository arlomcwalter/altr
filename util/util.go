package util

import (
	"altr/config"
	"github.com/mattn/go-runewidth"
)

// Math

func StrWidth(str string) int {
	var width int
	for _, char := range str {
		width += CharWidth(char)
	}
	return width
}

func CharWidth(char rune) int {
	if char == '\t' {
		return config.TabWidth
	} else {
		return runewidth.RuneWidth(char)
	}
}

func Min(min, val int) int {
	if val > min {
		return min
	}

	return val
}

//func Max(max, val int) int {
//	if val < max {
//		return max
//	}
//
//	return val
//}

func Clamp(min, max, val int) int {
	if val <= min {
		return min
	}

	if val >= max {
		return max
	}

	return val
}

//func Abs(val int) int {
//	if val > 0 {
//		return val
//	}
//
//	return int(math.Abs(float64(val)))
//}
