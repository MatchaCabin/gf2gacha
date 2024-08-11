package logic

import (
	"gf2gacha/model"
	"gf2gacha/preload"
	"github.com/pkg/errors"
)

func GetPoolInfo(uid string, poolType int64) (model.Pool, error) {
	localRecordList, err := GetLocalRecord(uid, poolType)
	if err != nil {
		return model.Pool{}, errors.WithStack(err)
	}

	pool := model.Pool{PoolType: poolType}
	for _, storedRecord := range localRecordList {
		pool.GachaCount++
		pool.StoredCount++
		item := preload.ItemMap[storedRecord.ItemId]
		if item.Rank == 5 {
			//检测是否歪
			var lose bool
			if upItemId, hasUp := preload.UpItemMap[storedRecord.PoolId]; hasUp && upItemId != storedRecord.ItemId {
				lose = true
			}

			pool.RecordList = append(pool.RecordList, model.DisplayRecord{
				Id:    item.Id,
				Name:  preload.LangMap[item.Name.Id],
				Lose:  lose,
				Count: pool.StoredCount,
			})

			pool.StoredCount = 0
			pool.Rank5Count++
		} else if item.Rank == 4 {
			pool.Rank4Count++
		} else {
			pool.Rank3Count++
		}
	}

	return pool, nil
}
