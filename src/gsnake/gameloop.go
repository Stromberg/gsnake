package main

import (
	"gsnake/gsnake"

	"time"

	termbox "github.com/nsf/termbox-go"
)

type gameloop struct {
	game *gsnake.Game
}

type keyType int

const (
	Move keyType = iota
	End
)

type keyEvent struct {
	command keyType
	dir     gsnake.Direction
}

func newGameloop() *gameloop {
	return &gameloop{
		game: gsnake.NewGame(20, 20),
	}
}

func (l *gameloop) start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	evChan := make(chan keyEvent)
	go l.listenToKeyboard(evChan)

	l.render()

mainloop:

	for {
		select {
		case e := <-evChan:
			switch e.command {
			case Move:
				l.game.Dir = e.dir
			default:
				break mainloop
			}
		default:
			l.game.Update()
			l.render()
			time.Sleep(250 * time.Millisecond)
		}
	}

}

func (l *gameloop) listenToKeyboard(evChan chan keyEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				evChan <- keyEvent{Move, gsnake.Left}
			case termbox.KeyArrowDown:
				evChan <- keyEvent{Move, gsnake.Down}
			case termbox.KeyArrowRight:
				evChan <- keyEvent{Move, gsnake.Right}
			case termbox.KeyArrowUp:
				evChan <- keyEvent{Move, gsnake.Up}
			case termbox.KeyEsc:
				evChan <- keyEvent{End, gsnake.Up}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (l *gameloop) render() {
	r := newGameRenderer(l.game)
	r.render()
}
