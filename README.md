# multiplayer-game

## Server architecture

```text
Client <-> WebSocket <-> Game Server <-> Database (optional)
              │
              ├── Auth Service
              ├── Matchmaking
              ├── Game Logic
              ├── State Management
              └── Broadcast System
```

## Project structure

```text
/game-server
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── game/
│   ├── websocket/
│   ├── matchmaking/
│   ├── models/
│   └── utils/
├── pkg/
│   └── protobuf/
└── config/
    └── config.yaml
```
