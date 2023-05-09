package gate

import (
	"context"

	"github.com/fengqk/mars-base/actor"
	"github.com/fengqk/mars-base/rpc"
)

type (
	IEventProcess interface {
		actor.IActor
	}

	EventProcess struct {
		actor.Actor
	}
)

func (e *EventProcess) Init() {
	e.Actor.Init()
	actor.MGR.RegisterActor(e)
	e.Actor.Start()
}

func (e *EventProcess) G_Player_Login(ctx context.Context, socketId uint32, mailbox rpc.MailBox) {
	actor.MGR.SendMsg(rpc.RpcHead{}, "PlayerMgr.ADD_ACCOUNT", socketId, mailbox)
}
