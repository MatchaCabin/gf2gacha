package logic

import (
	"gf2gacha/model"
	"gf2gacha/preload"
	"github.com/pkg/errors"
)

func GetPoolInfo(uid string, poolType int64) (model.Pool, error) {
	localRecordList, err := GetLocalRecord(uid, poolType, 0)
	if err != nil {
		return model.Pool{}, errors.WithStack(err)
	}

	pool := model.Pool{PoolType: poolType}
	var isPreviousLose bool
	for _, storedRecord := range localRecordList {
		pool.GachaCount++
		pool.StoredCount++
		item := preload.ItemMap[storedRecord.ItemId]
		itemRank := preload.ItemRankMap[storedRecord.PoolId][storedRecord.ItemId]
		if itemRank == 5 {
			if isPreviousLose {
				pool.GuaranteesCount++
			}
			//检测是否歪
			var lose bool
			if upItemId, hasUp := preload.UpItemMap[storedRecord.PoolId]; hasUp && upItemId != storedRecord.ItemId {
				pool.LoseCount++
				lose = true
				isPreviousLose = true
			} else {
				isPreviousLose = false
			}

			pool.RecordList = append(pool.RecordList, model.DisplayRecord{
				Id:    item.Id,
				Name:  preload.LangMap[item.Name.Id],
				Lose:  lose,
				Count: pool.StoredCount,
			})

			pool.StoredCount = 0
			pool.Rank5Count++
		} else if itemRank == 4 {
			pool.Rank4Count++
		} else {
			pool.Rank3Count++
		}
	}

	return pool, nil
}
