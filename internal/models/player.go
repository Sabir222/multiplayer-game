package models

import "github.com/Sabir222/multiplayer-game/internal/utils"

type Player struct {
	ID       string  `json:"id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Health   int     `json:"health"`
	Rotation float64 `json:"rotation"`
}

func NewPlayer() *Player {
	return &Player{
		ID:     utils.GenerateUUID(),
		Health: 100,
	}
}
