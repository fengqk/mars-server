package db

import (
	"context"
	
	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-base/db"
	"github.com/fengqk/mars-server/orm"
)

// 自动生成代码

func (this *Player) __LoadPlayerBaseDataDB(PlayerId int64) error{
    data := &orm.PlayerBaseData{PlayerId:PlayerId}
    rows, err := db.DB.Query(db.LoadSql(data, db.WithWhere(data)))
    rs, err := db.Query(rows, err)
    if err == nil && rs.Next() {
        db.LoadObjSql(&this.PlayerBaseData, rs.Row())
    }
	return err
}


func (this *Player) __SavePlayerBaseDataDB(){
	if this.PlayerBaseData.Dirty{
    	db.DB.Exec(db.SaveSql(this.PlayerBaseData))
		this.PlayerBaseData.Dirty = false
	}
}

func (this *Player) __SavePlayerBaseData(data orm.PlayerBaseData){
    this.PlayerBaseData = data
	this.PlayerBaseData.Dirty = true
    base.LOG.Printf("玩家[%d] SavePlayerBaseData", this.PlayerBaseData.PlayerId)
}

func (this *PlayerMgr) SavePlayerBaseData(ctx context.Context, playerId int64, data orm.PlayerBaseData){
    player := this.GetPlayer(playerId)
	if player != nil{
        player.__SavePlayerBaseData(data)
	}
}

func (this *Player) __LoadItemDataDB(PlayerId int64) error{
    data := &orm.ItemData{PlayerId:PlayerId}
    rows, err := db.DB.Query(db.LoadSql(data, db.WithWhere(data)))
    rs, err := db.Query(rows, err)
    if err == nil && rs.Next() {
        db.LoadObjSql(&this.ItemData, rs.Row())
    }
	return err
}


func (this *Player) __SaveItemDataDB(){
	if this.ItemData.Dirty{
    	db.DB.Exec(db.SaveSql(this.ItemData))
		this.ItemData.Dirty = false
	}
}

func (this *Player) __SaveItemData(data orm.ItemData){
    this.ItemData = data
	this.ItemData.Dirty = true
    base.LOG.Printf("玩家[%d] SaveItemData", this.ItemData.PlayerId)
}

func (this *PlayerMgr) SaveItemData(ctx context.Context, playerId int64, data orm.ItemData){
    player := this.GetPlayer(playerId)
	if player != nil{
        player.__SaveItemData(data)
	}
}

func (this *Player) LoadPlayerDB(PlayerId int64) error{
    this.Init(PlayerId)
    if err := this.__LoadPlayerBaseDataDB(PlayerId); err != nil{
        base.LOG.Printf("__LoadPlayerBaseDataDB() error")
        return err 
    }
    if err := this.__LoadItemDataDB(PlayerId); err != nil{
        base.LOG.Printf("__LoadItemDataDB() error")
        return err 
    }
    return nil
}


func (this *Player) SavePlayerDB(){
    this.__SavePlayerBaseDataDB()
    this.__SaveItemDataDB()
}

