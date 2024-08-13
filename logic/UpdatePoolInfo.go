package logic

import (
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
)

// incrementalUpdatePoolInfo 增量更新
func incrementalUpdatePoolInfo(logInfo model.LogInfo, poolType int64) error {
	localRecordList, err := GetLocalRecord(logInfo.Uid, poolType)
	if err != nil {
		return errors.WithStack(err)
	}

	var lastLocalRecord model.LocalRecord

	if len(localRecordList) > 0 {
		lastLocalRecord = localRecordList[len(localRecordList)-1]
	}

	var diffRemoteRecordList []model.RemoteRecord
	respData, err := FetchRemoteData(logInfo.GachaUrl, logInfo.AccessToken, "", poolType)
	if err != nil {
		return errors.WithStack(err)
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
			return errors.WithStack(err)
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
			return errors.WithStack(err)
		}
	}

	return nil
}

func IncrementalUpdatePoolInfo() (string, error) {
	logInfo, err := util.GetLogInfo()
	if err != nil {
		return "", errors.WithStack(err)
	}

	err = incrementalUpdatePoolInfo(logInfo, 1)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(logInfo, 3)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(logInfo, 4)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(logInfo, 5)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(logInfo, 8)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return logInfo.Uid, nil
}
