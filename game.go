package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mattn/go-tty"

	"github.com/thzoid/ws-game-server/shared"
)

var (
	server        *websocket.Conn
	matchMap      *shared.Map
	players       map[uuid.UUID]shared.Player
	localPlayerID uuid.UUID
)

func renderMap() {
	// Check if map exists
	if matchMap == nil {
		return
	}
	// Prepare player list
	playersToRender := make([]uuid.UUID, 0, len(players))
	for k := range players {
		playersToRender = append(playersToRender, k)
	}

	// Render map
	tm.MoveCursor(1, 1)
	for i := 0; i < matchMap.Size.Y; i++ {
		for j := 0; j < matchMap.Size.X; j++ {
			char := '.'
			for k, p := range playersToRender {
				if players[p].Position.Equals(shared.Coordinate{X: j, Y: i}) {
					char = players[p].UserProfile.Nick
					// Delete player from slice
					playersToRender[k] = playersToRender[len(playersToRender)-1]
					playersToRender = playersToRender[:len(playersToRender)-1]
					break
				}
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

func readInput() {
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

			// Move player in server
			shared.WriteMessage(server, "move", shared.MoveRequest{Direction: direction})

			// Move player locally
			if player, ok := players[localPlayerID]; ok {
				player.Move(direction, *matchMap)
				players[localPlayerID] = player
			}
		}
	}()
}
