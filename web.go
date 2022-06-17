package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/thzoid/ws-game-server/shared"
)

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("server disconnected")
			return
		}

		fmt.Println("message received:", string(p))
	}
}

func connect(url string, hs shared.CtS_HandshakeRequest) {
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		fmt.Println("could not connect to server")
		return
	}
	fmt.Println("connected to server")

	// Send hanshake to server
	body, _ := json.Marshal(hs)
	req := shared.Request{
		Type: "hanshake",
		Body: body,
	}
	conn.WriteJSON(req)

	reader(conn)
}
