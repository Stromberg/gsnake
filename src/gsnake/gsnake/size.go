package gsnake

type size struct {
	width  int
	height int
}

func (s size) isOutside(pos Position) bool {
	return pos.X < 0 || pos.X >= s.width || pos.Y < 0 || pos.Y >= s.height
}
