package gsnake

type Position struct {
	X int
	Y int
}

func (p Position) isSame(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Position) offset(dx, dy int) Position {
	return Position{p.X + dx, p.Y + dy}
}
