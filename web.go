package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

// Input function
func reader(conn *websocket.Conn) {
	for {
		m, err := shared.ReadMessage(conn)
		if err != nil {
			fmt.Println("server disconnected")
			os.Exit(0)
		}

		switch m.Type {
		case "handshake":
			// Read handshake from server
			hsS := new(shared.HandshakeResponse)
			json.Unmarshal(m.Body, hsS)

			// Set up map size
			matchMap = new(shared.Map)
			*matchMap = hsS.MatchMap

			// Get Player ID
			localPlayerID = hsS.PlayerID
		case "heartbeat":
			// Read heartbeat from server
			hb := new(shared.HeartbeatResponse)
			json.Unmarshal(m.Body, hb)

			// Update player list locally
			players = hb.Players
		default:
			fmt.Println("message received:", m)
		}
	}
}

func connect(url string, hs shared.HandshakeRequest) {
	// Connect to server
	var err error
	server, _, err = websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		fmt.Println("could not connect to server")
		os.Exit(0)
	}
	fmt.Println("connected to server")

	// Send hanshake to server
	shared.WriteMessage(server, "handshake", hs)

	// Read input from user
	readInput()

	// Start reading messages from server
	reader(server)
}
