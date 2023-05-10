package orm

type (
	DirtyData struct {
		Dirty bool `sql:"-"`
	}

	PlayerBaseData struct {
		DirtyData
		table          string `sql:"table;name:tb_players"`
		PlayerId       int64  `sql:"primary;name:charguid"`
		AccountId      int64  `sql:"name:account_id"`
		PlayerName     string `sql:"name:name"`
		Level          int    `sql:"name:level"`
		LastLogoutTime int64  `sql:"datetime;name:last_logout_time"`
		LastLoginTime  int64  `sql:"datetime;name:last_login_time"`
	}

	ItemData struct {
		DirtyData
		table    string `sql:"table;name:tb_items"`
		PlayerId int64  `sql:"primary;name:charguid"`
		Bag      Bag    `sql:"blob;name:data"`
	}

	PlayerData struct {
		PlayerBaseData
		ItemData
	}
)

func (this *PlayerData) Init(PlayerId int64) {
	this.ItemData = ItemData{PlayerId: PlayerId, Bag: Bag{DataMap: map[int64]*Item{}}}
}
