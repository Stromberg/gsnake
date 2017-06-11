package gsnake

type food struct {
	pos Position
}

type FoodRenderer interface {
	DrawFood(pos Position)
}

func newFood() *food {
	return &food{Position{0, 0}}
}

func (f *food) render(r FoodRenderer) {
	r.DrawFood(f.pos)
}
