package websocket

import (
	"net/http"

	"github.com/Sabir222/multiplayer-game/internal/models"
)

func HandleConnections(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		Player: models.NewPlayer(),
	}

	hub.register <- client
	go client.ReadPump()
	go client.WritePump()
}
