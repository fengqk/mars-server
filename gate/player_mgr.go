package gate

import (
	"context"
	"sync"

	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/rpc"
)

type (
	IPlayerMgr interface {
		actor.IActor
		AddPlayerMap(uint32, rpc.MailBox) int
		ReleaseSocketMap(uint32, bool)
		GetSocket(int64) uint32
		GetPlayerId(uint32) int64
		GetPlayer(uint32) *Player
	}

	PlayerMgr struct {
		actor.Actor
		socketMap map[uint32]int64
		playerMap map[int64]*Player
		locker    *sync.RWMutex
	}
)

func (p *PlayerMgr) AddPlayerMap(socketId uint32, mailbox rpc.MailBox) int {
	playerId := mailbox.Id
	Id := p.GetSocket(playerId)
	p.ReleaseSocketMap(Id, Id != socketId)

	player := NewPlayer(socketId, playerId)
	player.GClusterId = mailbox.ClusterId
	p.locker.Lock()
	p.playerMap[playerId] = player
	p.socketMap[socketId] = playerId
	p.locker.Unlock()
	SERVER.GetCluster().SendMsg(rpc.RpcHead{ClusterId: player.GClusterId, Id: playerId}, "game<-PlayerMgr.Player_Login", SERVER.GetCluster().Id(), mailbox)
	return base.NONE_ERROR
}

func (p *PlayerMgr) ReleaseSocketMap(socketId uint32, bClose bool) {
	p.locker.RLock()
	playerId, _ := p.socketMap[socketId]
	p.locker.RUnlock()
	p.locker.Lock()
	delete(p.playerMap, playerId)
	delete(p.socketMap, socketId)
	p.locker.Unlock()
	SERVER.GetServer().StopClient(socketId)
}

func (p *PlayerMgr) GetSocket(playerId int64) uint32 {
	socketId := uint32(0)
	p.locker.RLock()
	player, exist := p.playerMap[playerId]
	p.locker.RUnlock()
	if exist {
		socketId = player.SocketId
	}
	return socketId
}

func (p *PlayerMgr) GetPlayerId(socketId uint32) int64 {
	playerId := int64(0)
	p.locker.RLock()
	id, exist := p.socketMap[socketId]
	p.locker.RUnlock()
	if exist {
		playerId = id
	}
	return playerId
}

func (p *PlayerMgr) GetPlayer(socketId uint32) *Player {
	playerId := p.GetPlayerId(socketId)
	p.locker.RLock()
	player, exist := p.playerMap[playerId]
	p.locker.RUnlock()
	if exist {
		return player
	}
	return nil
}

func (p *PlayerMgr) Init() {
	p.Actor.Init()
	p.socketMap = make(map[uint32]int64)
	p.playerMap = make(map[int64]*Player)
	p.locker = &sync.RWMutex{}
	actor.MGR.RegisterActor(p)
	p.Actor.Start()
}

func (p *PlayerMgr) ADD_ACCOUNT(ctx context.Context, socketId uint32, mailbox rpc.MailBox) {
	base.LOG.Printf("login incoming  Socket:%d PlayerId:%d GClusterId:%d ", socketId, mailbox.Id, mailbox.ClusterId)
	p.AddPlayerMap(socketId, mailbox)
}

func (p *PlayerMgr) DEL_ACCOUNT(ctx context.Context, socketid uint32) {
	playerId := p.GetPlayerId(socketid)
	p.ReleaseSocketMap(socketid, true)
	SERVER.GetCluster().SendMsg(rpc.RpcHead{SendType: rpc.SEND_BOARD_CAST}, "game<-PlayerMgr.G_ClientLost", playerId)
}
