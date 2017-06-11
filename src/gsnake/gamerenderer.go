package main

import (
	"fmt"
	"gsnake/gsnake"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

type gameRenderer struct {
	game                     *gsnake.Game
	top, bottom, left, right int
	defaultColor             termbox.Attribute
	bgColor                  termbox.Attribute
	snakeColor               termbox.Attribute
	foodColor                termbox.Attribute
}

func newGameRenderer(game *gsnake.Game) *gameRenderer {
	w, h := termbox.Size()
	midY := h / 2
	left := (w - game.Width) / 2
	right := (w + game.Width) / 2
	top := midY - (game.Height / 2)
	bottom := midY + (game.Height / 2) + 1

	return &gameRenderer{
		game:         game,
		left:         left,
		right:        right,
		top:          top,
		bottom:       bottom,
		defaultColor: termbox.ColorDefault,
		bgColor:      termbox.ColorDefault,
		snakeColor:   termbox.ColorYellow,
		foodColor:    termbox.ColorRed,
	}
}

func (r *gameRenderer) render() {
	termbox.Clear(r.defaultColor, r.bgColor)

	r.drawBox()
	r.game.Render(r)

	_ = termbox.Flush()
}

func (r *gameRenderer) DrawScore(score, bestScore int) {
	msg := fmt.Sprintf("Score: %v (%v)", score, bestScore)
	r.tbprint(r.left, r.bottom+1, r.defaultColor, r.defaultColor, msg)
}

func (r *gameRenderer) DrawFood(pos gsnake.Position) {
	termbox.SetCell(r.left+pos.X, r.top+pos.Y+1, ' ', r.foodColor, r.foodColor)
}

func (r *gameRenderer) DrawBodyPart(pos gsnake.Position) {
	termbox.SetCell(r.left+pos.X, r.top+pos.Y+1, ' ', r.snakeColor, r.snakeColor)
}

func (r *gameRenderer) drawBox() {
	for i := r.top; i < r.bottom; i++ {
		termbox.SetCell(r.left-1, i, '│', r.defaultColor, r.bgColor)
		termbox.SetCell(r.left+r.game.Width, i, '│', r.defaultColor, r.bgColor)
	}

	termbox.SetCell(r.left-1, r.top, '┌', r.defaultColor, r.bgColor)
	termbox.SetCell(r.left-1, r.bottom, '└', r.defaultColor, r.bgColor)
	termbox.SetCell(r.left+r.game.Width, r.top, '┐', r.defaultColor, r.bgColor)
	termbox.SetCell(r.left+r.game.Width, r.bottom, '┘', r.defaultColor, r.bgColor)

	r.fill(r.left, r.top, r.game.Width, 1, termbox.Cell{Ch: '─'})
	r.fill(r.left, r.bottom, r.game.Width, 1, termbox.Cell{Ch: '─'})
}

func (r *gameRenderer) fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func (r *gameRenderer) tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
