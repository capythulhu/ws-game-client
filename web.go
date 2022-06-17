package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("server disconnected")
			os.Exit(0)
		}

		r := &shared.Request{}
		json.Unmarshal(p, r)
		switch r.Type {
		case "handshake":
			// Read handshake from server
			hsS := &shared.HandshakeResponse{}
			json.Unmarshal(r.Body, hsS)
			fmt.Println("handshake received.", "server map:", hsS.MatchMap)

			// Set up map size
			matchMap = new(shared.Map)
			*matchMap = hsS.MatchMap
		case "move":

		default:
			fmt.Println("message received:", string(p))
		}
	}
}

func connect(url string, hs shared.HandshakeRequest) {
	// Connect to server
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		fmt.Println("could not connect to server")
		os.Exit(0)
	}
	fmt.Println("connected to server")

	// Send hanshake to server
	shared.WriteRequest(conn, "handshake", hs)

	// Start reading messages from server
	reader(conn)
}
