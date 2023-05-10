package player

import (
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/cluster"
	"github.com/fengqk/mars-base/rpc"
)

// 自动生成代码

func (this *Player) SavePlayerBaseData(){
	this.PlayerBaseData.Dirty = true
}

func (this *Player) __SavePlayerBaseDataDB(){
	if this.PlayerBaseData.Dirty{
    	cluster.MGR.SendMsg(rpc.RpcHead{DestServerType:rpc.SERVICE_DB, Id:this.MailBox.Id}, "PlayerMgr.SavePlayerBaseData", this.PlayerBaseData.PlayerId, this.PlayerBaseData)
		this.PlayerBaseData.Dirty = false
    	base.LOG.Printf("玩家[%d] SavePlayerBaseData", this.MailBox.Id)
	}
}

func (this *Player) SaveItemData(){
	this.ItemData.Dirty = true
}

func (this *Player) __SaveItemDataDB(){
	if this.ItemData.Dirty{
    	cluster.MGR.SendMsg(rpc.RpcHead{DestServerType:rpc.SERVICE_DB, Id:this.MailBox.Id}, "PlayerMgr.SaveItemData", this.ItemData.PlayerId, this.ItemData)
		this.ItemData.Dirty = false
    	base.LOG.Printf("玩家[%d] SaveItemData", this.MailBox.Id)
	}
}

func (this *Player) SavePlayerDB(){
    this.__SavePlayerBaseDataDB()
    this.__SaveItemDataDB()
}

