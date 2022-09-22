package util

import (
	"bufio"
	"fmt"
	"os"
)

func Parse(path string) []string {
	file, err := os.Stat(path)

	if err == nil && !file.IsDir() {
		content, err := read(path)
		if err == nil {
			return content
		}
	}

	return []string{""}
}

func read(path string) ([]string, error) {
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

func Write(lines []string, path string) error {
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
