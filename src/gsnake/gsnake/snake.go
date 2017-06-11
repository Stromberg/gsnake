package gsnake

type snake struct {
	parts           []bodyPart
	dir             Direction
	pendingSegments int
}

type bodyPart struct {
	pos Position
	dir Direction
}

type SnakeRenderer interface {
	DrawBodyPart(pos Position)
}

func newSnake() *snake {
	s := &snake{}
	s.reset()
	return s
}

func (s *snake) reset() {
	s.parts = []bodyPart{
		bodyPart{Position{5, 0}, Right},
		bodyPart{Position{4, 0}, Right},
		bodyPart{Position{3, 0}, Right},
	}

	s.dir = Right
	s.pendingSegments = 0
}

func (s *snake) changeDirection(dir Direction) {
	if s.dir == Left && dir == Right {
		return
	}

	if s.dir == Right && dir == Left {
		return
	}

	if s.dir == Up && dir == Down {
		return
	}

	if s.dir == Down && dir == Up {
		return
	}

	s.dir = dir
}

func (s *snake) render(renderer SnakeRenderer) {
	for _, p := range s.parts {
		renderer.DrawBodyPart(p.pos)
	}
}

func (s *snake) intersectsGrid(gridSize size) bool {
	for _, p := range s.parts {
		if gridSize.isOutside(p.pos) {
			return true
		}
	}

	return false
}

func (s *snake) intersects(pos Position) bool {
	for _, p := range s.parts {
		if pos.isSame(p.pos) {
			return true
		}
	}

	return false
}

func (s *snake) move() {
	last := s.parts[len(s.parts)-1]
	newPart := bodyPart{last.pos, last.dir}

	s.parts[0].dir = s.dir
	for i := len(s.parts) - 1; i >= 0; i-- {
		switch s.parts[i].dir {
		case Left:
			s.parts[i].pos = s.parts[i].pos.offset(-1, 0)
		case Right:
			s.parts[i].pos = s.parts[i].pos.offset(1, 0)
		case Up:
			s.parts[i].pos = s.parts[i].pos.offset(0, -1)
		case Down:
			s.parts[i].pos = s.parts[i].pos.offset(0, 1)
		}

		if i > 0 {
			s.parts[i].dir = s.parts[i-1].dir
		}
	}

	if s.pendingSegments > 0 {
		s.parts = append(s.parts, newPart)
		s.pendingSegments--
	}
}

func (s *snake) addBodySegment() {
	s.pendingSegments++
}

func (s *snake) isSelfIntersecting() bool {
	for i := range s.parts {
		for j := range s.parts {
			if i != j {
				p1 := s.parts[i]
				p2 := s.parts[j]
				if p1.pos.isSame(p2.pos) {
					return true
				}
			}
		}
	}

	return false
}
