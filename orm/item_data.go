package orm

type (
	Item struct {
		Id       int64
		PlayerId int64
		ItemId   int32
		Count    int32
	}

	Bag struct {
		DataMap map[int64]*Item
	}
)
