package window

import "github.com/gdamore/tcell/v2"

var (
	// Colors
	accent       = tcell.NewRGBColor(165, 102, 227)
	text         = tcell.NewRGBColor(255, 255, 255)
	accentDimmed = tcell.NewRGBColor(147, 116, 179)

	// Styles
	resetStyle     = tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	titleBarStyle  = tcell.StyleDefault.Foreground(text).Background(accent)
	statusBarStyle = tcell.StyleDefault.Foreground(text).Background(accentDimmed)

	lineNumStyle    = tcell.StyleDefault.Foreground(accentDimmed).Background(tcell.ColorReset)
	activeLineStyle = tcell.StyleDefault.Foreground(text).Background(tcell.ColorReset)
)
