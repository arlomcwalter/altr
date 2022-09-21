package editor

type Cursor struct {
	X, Y int
}

func (c *Cursor) MoveX(move int, e *Editor) {
	c.SetX(c.X+move, e)
}

func (c *Cursor) MoveY(move int, e *Editor) {
	c.SetY(c.Y+move, e)
}

func (c *Cursor) Set(x, y int, e *Editor) {
	c.SetX(x, e)
	c.SetY(y, e)
}

func (c *Cursor) SetX(x int, e *Editor) {
	c.X = clamp(0, e.getWidth(), x)
}

func (c *Cursor) SetY(y int, e *Editor) {
	c.Y = clamp(1, e.getHeight(), y)
}
