package editor

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
)

func drawStr(msg string, x, y int, fg, bg termbox.Attribute) int {
	for _, char := range msg {
		x += drawChar(char, x, y, fg, bg)
	}

	return x
}

func drawChar(char rune, x, y int, fg, bg termbox.Attribute) int {
	termbox.SetCell(x, y, char, fg, bg)
	return runewidth.RuneWidth(char)
}

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

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func (e *Editor) docWidth() int {
	var longest int

	for _, str := range e.content {
		length := strLen(str)
		if length > longest {
			longest = length
		}
	}

	termWidth, _ := termbox.Size()
	return max(termWidth, longest)
}

func (e *Editor) docHeight() int {
	_, termHeight := termbox.Size()
	return max(termHeight-1, len(e.content))
}

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
