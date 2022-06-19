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
			hsS := &shared.HandshakeResponse{}
			json.Unmarshal(m.Body, hsS)

			// Set up map size
			matchMap = new(shared.Map)
			*matchMap = hsS.MatchMap
		case "heartbeat":
			// Read heartbeat from server
			hb := &shared.HeartbeatResponse{}
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
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		fmt.Println("could not connect to server")
		os.Exit(0)
	}
	fmt.Println("connected to server")

	// Send hanshake to server
	shared.WriteMessage(conn, "handshake", hs)

	// Start reading messages from server
	reader(conn)
}
