package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/mattn/go-tty"

	"github.com/thzoid/ws-game-server/shared"
)

func renderMap(size shared.Coordinate) {
	tm.MoveCursor(1, 1)
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			var char rune
			if localPlayer.Actor.Position.Equals(shared.Coordinate{X: j, Y: i}) {
				char = localPlayer.nick
			} else {
				char = '.'
			}
			fmt.Printf(string(char) + " ")
		}
		fmt.Printf("\n")
	}
	tm.Flush()
}

func render() {
	tm.Clear()
	for range time.Tick(time.Millisecond * 100) {
		renderMap(shared.Coordinate{X: mapWidth, Y: mapHeight})
	}
}

func read() {
	go func() {
		tty, _ := tty.Open()
		defer tty.Close()
		for range time.Tick(time.Millisecond * 100) {
			r, _ := tty.ReadRune()
			direction := shared.Coordinate{X: 0, Y: 0}
			switch r {
			case 'w':
				direction.Y = -1
			case 'a':
				direction.X = -1
			case 's':
				direction.Y = 1
			case 'd':
				direction.X = 1
			}

			localPlayer.Move(direction)
		}
	}()
}
