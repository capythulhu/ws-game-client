package main

import "flag"

const (
	mapWidth  = 8
	mapHeight = 30
)

var localPlayer = Player{
	Actor: Actor{
		position: Coordinate{3, 2},
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

	connect(*urlPtr, CtS_HandshakeRequest{Nick: rune((*nickPtr)[0])})
}
