package model

import "time"

type RemoteRecord struct {
	PoolId         int64 `json:"pool_id"`
	ItemId         int64 `json:"item"`
	GachaTimestamp int64 `json:"time"`
}

type LocalRecord struct {
	Id             int64
	InsertTime     time.Time `xorm:"created"`
	PoolType       int64
	PoolId         int64
	ItemId         int64
	GachaTimestamp int64
}

type DisplayRecord struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Lose  bool   `json:"lose"`
	Count int64  `json:"count"`
}
