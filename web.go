package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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

func connect(url string, hsObj CtS_HandshakeRequest) {
	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		fmt.Println("could not connect to server")
		return
	}
	fmt.Println("connected to server")

	// Send hanshake to server
	body, _ := json.Marshal(hsObj)
	req := Request{
		Type: "hanshake",
		Body: body,
	}
	conn.WriteJSON(req)

	reader(conn)
}
