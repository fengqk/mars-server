package player

import (
	"context"
	"reflect"

	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/cluster"
	"github.com/fengqk/mars-base/rpc"
)

var (
	MGR PlayerMgr
)

type (
	IPlayerMgr interface {
		Update()
	}

	PlayerMgr struct {
		actor.Actor
		actor.ActorPoolDynamic
	}
)

func (p *PlayerMgr) Init() {
	p.Actor.Init()
	p.InitActor(p, reflect.TypeOf(Player{}))
	actor.MGR.RegisterActor(p)
	p.Actor.Start()
}

// 玩家登录
func (p *PlayerMgr) LoginPlayerRequset(ctx context.Context, playerId int64, gateClusterId uint32, socketId uint32) {
	pMailBox := cluster.MGR.MailBox.Get(playerId)
	if pMailBox == nil {
		info := &rpc.MailBox{}
		info.Id = playerId
		info.ClusterId = cluster.MGR.Id()
		info.MailType = rpc.MAIL_Player
		if cluster.MGR.MailBox.Create(info) {
			pMailBox = info
		} else {
			return
		}
	}

	cluster.MGR.SendMsg(rpc.RpcHead{ClusterId: gateClusterId}, "gate<-EventProcess.G_Player_Login", socketId, *pMailBox)
}

// 玩家登录
func (p *PlayerMgr) Player_Login(ctx context.Context, gateClusterId uint32, mailbox rpc.MailBox) {
	playerId := mailbox.Id
	if p.GetActor(playerId) != nil {
		actor.MGR.SendMsg(rpc.RpcHead{Id: playerId}, "Player.ReLogin", gateClusterId, mailbox)
		return
	}

	pPlayer := &Player{}
	pPlayer.PlayerId = playerId
	pPlayer.SetId(playerId)
	pPlayer.Init()
	p.AddActor(pPlayer)
	actor.MGR.SendMsg(rpc.RpcHead{Id: playerId}, "Player.Login", gateClusterId, mailbox)
}

// 玩家断开链接
func (p *PlayerMgr) G_ClientLost(ctx context.Context, playerId int64) {
	actor.MGR.SendMsg(rpc.RpcHead{Id: playerId}, "Player.Logout", playerId)
}
