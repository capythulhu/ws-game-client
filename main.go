package main

import (
	"flag"

	"github.com/thzoid/ws-game-server/shared"
)

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

	go connect(*urlPtr, shared.HandshakeRequest{Nick: rune((*nickPtr)[0])})
	render()
}
