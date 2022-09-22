package main

import (
	"altr/window"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	var name string
	switch len(args) {
	case 0:
		name = "untitled"
	case 1:
		name = args[0]
	default:
		fmt.Println("Expected 0 or 1 arguments.")
		os.Exit(0)
	}

	w := window.Create(name)
	w.Init()
	defer w.Shutdown()
}
