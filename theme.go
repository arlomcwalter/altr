package main

import "github.com/gdamore/tcell/v2"

var (
	// Colors
	titleColor  = tcell.NewRGBColor(165, 102, 227)
	textColor   = tcell.NewRGBColor(255, 255, 255)
	statusColor = tcell.NewRGBColor(147, 116, 179)

	// Styles
	resetStyle     = tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	titleBarStyle  = tcell.StyleDefault.Foreground(textColor).Background(titleColor)
	statusBarStyle = tcell.StyleDefault.Foreground(textColor).Background(statusColor)

	lineNumStyle    = tcell.StyleDefault.Foreground(statusColor).Background(tcell.ColorReset)
	activeLineStyle = tcell.StyleDefault.Foreground(textColor).Background(tcell.ColorReset)
)
