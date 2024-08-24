package logic

import (
	"cmp"
	"fmt"
	"gf2gacha/logger"
	"gf2gacha/model"
	"slices"
	"strconv"
)

func MergeRecord(recordList1, recordList2 []model.LocalRecord) (newRecordList []model.LocalRecord) {
	recordList1 = addFakeId(recordList1)
	recordList2 = addFakeId(recordList2)
	tmpRecordList := append(recordList1, recordList2...)
	slices.SortStableFunc(tmpRecordList, func(a, b model.LocalRecord) int {
		return cmp.Compare(a.Id, b.Id)
	})

	fakeIdMap := make(map[int64]struct{})
	for _, record := range tmpRecordList {
		if _, exist := fakeIdMap[record.Id]; !exist {
			fakeIdMap[record.Id] = struct{}{}
			newRecordList = append(newRecordList, model.LocalRecord{
				PoolType:       record.PoolType,
				PoolId:         record.PoolId,
				ItemId:         record.ItemId,
				GachaTimestamp: record.GachaTimestamp,
			})
		}
	}

	return newRecordList
}

func addFakeId(recordList []model.LocalRecord) []model.LocalRecord {
	if len(recordList) == 0 {
		return nil
	}
	var preTimeStamp, order int64
	for i, record := range recordList {
		if record.GachaTimestamp != preTimeStamp {
			//前一个只能是0或9
			if order != 0 && order != 9 {
				logger.Logger.Warnf("存在抽卡时间相同但不足十连的数据:%+v", recordList[i-1])
			}
			order = 0
		} else {
			order++
			if order > 9 {
				logger.Logger.Warnf("存在抽卡时间相同但超过十连的数据:%+v", recordList[i-1])
			}
		}

		fakeId, _ := strconv.ParseInt(fmt.Sprintf("%d%02d", record.GachaTimestamp, order), 10, 64)
		recordList[i].Id = fakeId

		preTimeStamp = record.GachaTimestamp
	}
	return recordList
}
