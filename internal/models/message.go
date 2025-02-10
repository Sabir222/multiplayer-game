package models

type BaseMessage struct {
	Type string `json:"type"`
}

type PlayerMoveMessage struct {
	Type    string `json:"type"`
	Payload struct {
		X         float64 `json:"x"`
		Y         float64 `json:"y"`
		Timestamp int64   `json:"timestamp"`
	} `json:"payload"`
}
