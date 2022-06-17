package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/mattn/go-tty"

	"github.com/thzoid/ws-game-server/shared"
)

var (
	matchMap      *shared.Map
	players       map[string]shared.Player
	localPlayerID string
)

func renderMap() {
	if matchMap == nil {
		return
	}

	tm.MoveCursor(1, 1)
	for i := 0; i < matchMap.Size.Y; i++ {
		for j := 0; j < matchMap.Size.X; j++ {
			var char rune
			if player, ok := players[localPlayerID]; ok && player.Position.Equals(shared.Coordinate{X: i, Y: j}) {
				char = players[localPlayerID].Nick
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
		renderMap()
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

			player := players[localPlayerID]
			player.Move(direction)
			players[localPlayerID] = player
		}
	}()
}
