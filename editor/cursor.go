package editor

type Cursor struct {
	x, y int
}

func (c *Cursor) MoveX(move int, e *Editor) {
	c.SetX(c.x+move, e)
}

func (c *Cursor) MoveY(move int, e *Editor) {
	c.SetY(c.y+move, e)
}

func (c *Cursor) Set(x, y int, e *Editor) {
	c.SetX(x, e)
	c.SetY(y, e)
}

func (c *Cursor) SetX(x int, e *Editor) {
	c.x = clamp(0, e.docWidth(), x)
}

func (c *Cursor) SetY(y int, e *Editor) {
	c.y = clamp(1, e.docHeight()+1, y)
}
