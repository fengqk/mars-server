package player

import "time"

type (
	Player struct {
		PlayerId   int64
		LastTime   int64
		SocketId   uint32
		GClusterId uint32
		ZClusterId uint32
	}
)

var (
	g_pPlayer = &Player{}
)

func NewPlayer(socketId uint32, playerId int64) *Player {
	player := Player{PlayerId: playerId, LastTime: time.Now().Unix(), SocketId: socketId}
	return &player
}
