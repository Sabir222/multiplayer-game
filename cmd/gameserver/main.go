package main

import (
	"log"
	"net/http"

	"github.com/Sabir222/multiplayer-game/internal/game"
	"github.com/Sabir222/multiplayer-game/internal/websocket"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	go game.GameLoop(hub)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleConnections(hub, w, r)
	})

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
