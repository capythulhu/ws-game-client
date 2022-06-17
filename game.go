package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/mattn/go-tty"
)

func renderMap(size Coordinate) {
	tm.MoveCursor(1, 1)
	for i := 0; i < size.x; i++ {
		for j := 0; j < size.y; j++ {
			var char rune
			if localPlayer.position.Equals(Coordinate{j, i}) {
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
		renderMap(Coordinate{mapWidth, mapHeight})
	}
}

func read() {
	go func() {
		tty, _ := tty.Open()
		defer tty.Close()
		for range time.Tick(time.Millisecond * 100) {
			r, _ := tty.ReadRune()
			direction := Coordinate{x: 0, y: 0}
			switch r {
			case 'w':
				direction.y = -1
			case 'a':
				direction.x = -1
			case 's':
				direction.y = 1
			case 'd':
				direction.x = 1
			}

			localPlayer.Move(direction)
		}
	}()
}
