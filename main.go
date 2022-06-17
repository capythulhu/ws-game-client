package main

import (
	"flag"

	"github.com/thzoid/ws-game-server/shared"
)

const (
	mapWidth  = 8
	mapHeight = 30
)

var localPlayer = Player{
	Actor: shared.Actor{
		Position: shared.Coordinate{X: 3, Y: 2},
	},
	velocity: 1,
}

func main() {
	urlPtr := flag.String("url", "ws://localhost:8080/", "server url for the player to connect")
	nickPtr := flag.String("nick", "c", "player nick to be shown on the map")

	flag.Parse()
	if len(*nickPtr) > 1 {
		panic("player nick must be a single char")
	}
	nick := rune((*nickPtr)[0])
	if (nick < 'a' || nick > 'z') &&
		(nick < 'A' || nick > 'Z') &&
		(nick < '0' || nick > '9') {
		panic("player nick must be a letter or a number")
	}

	connect(*urlPtr, shared.CtS_HandshakeRequest{Nick: rune((*nickPtr)[0])})
}
