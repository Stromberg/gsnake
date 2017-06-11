package gsnake

import "math/rand"

type Game struct {
	Width     int
	Height    int
	snake     *snake
	food      *food
	score     int
	bestScore int
	Dir       Direction
}

type GameRenderer interface {
	SnakeRenderer
	FoodRenderer
	DrawScore(score, bestScore int)
}

func NewGame(width, height int) *Game {
	g := &Game{
		Width:     width,
		Height:    height,
		snake:     newSnake(),
		food:      newFood(),
		score:     0,
		bestScore: 0,
		Dir:       Right,
	}

	g.reset()

	return g
}

func (g *Game) reset() {
	g.snake.reset()
	if g.score > g.bestScore {
		g.bestScore = g.score
	}
	g.score = 0
	g.Dir = Right
	g.placeFood()
}

func (g *Game) placeFood() {
	g.food.pos.X = rand.Intn(g.Width)
	g.food.pos.Y = rand.Intn(g.Height)
}

func (g *Game) Render(renderer GameRenderer) {
	g.snake.render(renderer)
	g.food.render(renderer)
	renderer.DrawScore(g.score, g.bestScore)
}

func (g *Game) Update() {
	g.snake.changeDirection(g.Dir)
	g.snake.move()
	g.checkForCollisions()
}

func (g *Game) checkForCollisions() {
	if g.snake.intersectsGrid(size{g.Width, g.Height}) || g.snake.isSelfIntersecting() {
		g.reset()
		return
	}

	if g.snake.intersects(g.food.pos) {
		g.score++
		g.placeFood()
		g.snake.addBodySegment()
	}
}
