package game

import (
	"time"

	"github.com/Sabir222/multiplayer-game/internal/websocket"
)

const TICK_RATE = 60

func GameLoop(hub *websocket.Hub) {
	tickInterval := time.Second / TICK_RATE
	ticker := time.NewTicker(tickInterval)

	for {
		select {
		case <-ticker.C:
			hub.mutex.RLock()
			for _, gameState := range hub.gameStates {
				gameState.Mutex.Lock()
				// Update game logic here
				gameState.Mutex.Unlock()
				gameState.BroadcastState()
			}
			hub.mutex.RUnlock()
		}
	}
}
