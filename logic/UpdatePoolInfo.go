package logic

import (
	"fmt"
	"gf2gacha/logger"
	"gf2gacha/model"
	"gf2gacha/preload"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"slices"
)

func UpdatePoolInfo(isFull bool) (messageList []string, err error) {
	logInfo, err := util.GetLogInfo()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	messageList = append(messageList, logInfo.Uid)
	for _, poolTypeUnit := range preload.PoolTypeMap {
		if isFull {
			n, err := fullUpdatePoolInfo(logInfo, poolTypeUnit.Id)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			messageList = append(messageList, fmt.Sprintf("%s 全量更新%d条数据", poolTypeUnit.Name, n))
		} else {
			n, err := incrementalUpdatePoolInfo(logInfo, poolTypeUnit.Id)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			messageList = append(messageList, fmt.Sprintf("%s 增量更新%d条数据", poolTypeUnit.Name, n))
		}
	}

	return messageList, nil
}

// incrementalUpdatePoolInfo 增量更新
func incrementalUpdatePoolInfo(logInfo model.LogInfo, poolType int64) (int, error) {
	localRecordList, err := GetLocalRecord(logInfo.Uid, poolType, 0)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	var lastLocalRecord model.LocalRecord

	if len(localRecordList) > 0 {
		lastLocalRecord = localRecordList[len(localRecordList)-1]
	}

	var diffRemoteRecordList []model.RemoteRecord
	respData, err := FetchRemoteData(logInfo.GachaUrl, logInfo.AccessToken, "", poolType)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	//todo 只对比一条是有可能有问题的
	var flag bool
	for i, remoteRecord := range respData.RecordList {
		if remoteRecord.ItemId == lastLocalRecord.ItemId && remoteRecord.GachaTimestamp == lastLocalRecord.GachaTimestamp {
			flag = true
			break
		} else {
			diffRemoteRecordList = append(diffRemoteRecordList, respData.RecordList[i])
		}
	}
	for respData.Next != "" && !flag {
		//time.Sleep(50 * time.Millisecond) //这个接口似乎没有限制频率
		respData, err = FetchRemoteData(logInfo.GachaUrl, logInfo.AccessToken, respData.Next, poolType)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		for i, remoteRecord := range respData.RecordList {
			if remoteRecord.ItemId == lastLocalRecord.ItemId && remoteRecord.GachaTimestamp == lastLocalRecord.GachaTimestamp {
				flag = true
				break
			} else {
				diffRemoteRecordList = append(diffRemoteRecordList, respData.RecordList[i])
			}
		}
	}

	if len(diffRemoteRecordList) > 0 {
		var diffLocalRecordList []model.LocalRecord
		for i := len(diffRemoteRecordList) - 1; i >= 0; i-- {
			diffLocalRecordList = append(diffLocalRecordList, model.LocalRecord{
				PoolType:       poolType,
				PoolId:         diffRemoteRecordList[i].PoolId,
				ItemId:         diffRemoteRecordList[i].ItemId,
				GachaTimestamp: diffRemoteRecordList[i].GachaTimestamp,
			})
		}
		err = SaveLocalRecord(logInfo.Uid, diffLocalRecordList)
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}

	updateNum := len(diffRemoteRecordList)
	logger.Logger.Infof("UID:%s poolType:%d 增量更新%d条数据", logInfo.Uid, poolType, updateNum)

	return updateNum, nil
}

// fullUpdatePoolInfo 全量更新
func fullUpdatePoolInfo(logInfo model.LogInfo, poolType int64) (int, error) {
	var remoteRecordList []model.LocalRecord

	respData, err := FetchRemoteData(logInfo.GachaUrl, logInfo.AccessToken, "", poolType)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	for _, record := range respData.RecordList {
		remoteRecordList = append(remoteRecordList, model.LocalRecord{
			PoolType:       poolType,
			PoolId:         record.PoolId,
			ItemId:         record.ItemId,
			GachaTimestamp: record.GachaTimestamp,
		})
	}

	for respData.Next != "" {
		//time.Sleep(50 * time.Millisecond) //这个接口似乎没有限制频率
		respData, err = FetchRemoteData(logInfo.GachaUrl, logInfo.AccessToken, respData.Next, poolType)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		for _, record := range respData.RecordList {
			remoteRecordList = append(remoteRecordList, model.LocalRecord{
				PoolType:       poolType,
				PoolId:         record.PoolId,
				ItemId:         record.ItemId,
				GachaTimestamp: record.GachaTimestamp,
			})
		}
	}

	if len(remoteRecordList) == 0 {
		return 0, nil
	}
	//合并前先备份
	err = util.BackupDB()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	slices.Reverse(remoteRecordList)

	localRecordList, err := GetLocalRecord(logInfo.Uid, poolType, remoteRecordList[0].GachaTimestamp)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	newRecordList := MergeRecord(remoteRecordList, localRecordList)

	err = RemoveLocalRecord(logInfo.Uid, poolType)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	err = SaveLocalRecord(logInfo.Uid, newRecordList)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	updateNum := len(newRecordList)
	logger.Logger.Infof("UID:%s poolType:%d 全量更新%d条数据", logInfo.Uid, poolType, updateNum)

	return updateNum, nil
}
