package game

import (
	"sync"
	"time"

	"github.com/Sabir222/multiplayer-game/internal/models"
)

type GameState struct {
	Players     map[string]*models.Player
	Projectiles map[string]*models.Projectile
	GameTime    time.Time
	Physics     *PhysicsEngine
	Mutex       sync.RWMutex
}

func NewGameState() *GameState {
	return &GameState{
		Players:     make(map[string]*models.Player),
		Projectiles: make(map[string]*models.Projectile),
		GameTime:    time.Now(),
		Physics:     NewPhysicsEngine(),
	}
}

func (gs *GameState) Update() {
	// Your state update logic
	gs.Mutex.Lock()
	defer gs.Mutex.Unlock()

	// Physics calculations
	gs.Physics.Step()

	// Update positions
	for _, p := range gs.Players {
		p.UpdatePosition()
	}

	// Check collisions
	gs.CheckCollisions()
}
