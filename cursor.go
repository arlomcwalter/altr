package main

import (
	"altr/util"
)

type Cursor struct {
	col, line int
}

// Absolute Position

func (c *Cursor) Clamp(cols, lines int, editor *Editor) {
	c.line = c.clampV(lines, editor)
	c.col = c.clampH(cols, editor)
}

func (c *Cursor) clampH(cols int, editor *Editor) int {
	maxVisibleI := cols - editor.padL - 1
	maxFileI := editor.calcWidthAt(c.line)
	maxEditorI := util.Min(maxVisibleI, maxFileI)

	clamped := util.Clamp(0, maxEditorI, c.col)

	if maxFileI >= maxEditorI {
		if c.col < 0 && editor.scrollX != 0 {
			editor.scrollX += c.col
		} else if c.col > maxVisibleI {
			editor.scrollX += c.col - maxEditorI
		}
	}

	return clamped
}

func (c *Cursor) clampV(lines int, editor *Editor) int {
	maxVisibleI := lines - editor.vPadding() - 1
	maxFileI := editor.lineCount() - 1
	maxEditorI := util.Min(maxVisibleI, maxFileI)

	clamped := util.Clamp(0, maxEditorI, c.line)

	if maxFileI > maxEditorI {
		if clamped+editor.scrollY == maxFileI {
			return clamped
		} else if c.line < 0 && editor.scrollY != 0 {
			editor.scrollY += c.line
		} else if c.line > maxVisibleI {
			editor.scrollY += c.line - maxEditorI
		}
	}

	return clamped
}

func (c *Cursor) Pos(editor *Editor) (int, int) {
	return editor.padL + c.col, editor.padT + c.line
}

// Relative Position

func (c *Cursor) MoveX(x int) {
	c.SetX(c.col + x)
}

func (c *Cursor) MoveY(y int) {
	c.SetY(c.line + y)
}

func (c *Cursor) Set(x, y int) {
	c.SetX(x)
	c.SetY(y)
}

func (c *Cursor) SetX(x int) {
	c.col = x
}

func (c *Cursor) SetY(y int) {
	c.line = y
}
