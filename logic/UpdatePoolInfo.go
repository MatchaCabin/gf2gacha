package logic

import (
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
)

// incrementalUpdatePoolInfo 增量更新
func incrementalUpdatePoolInfo(gameInfo model.Info, poolType int64) error {
	localRecordList, err := GetLocalRecord(gameInfo.Uid, poolType)
	if err != nil {
		return errors.WithStack(err)
	}

	var lastLocalRecord model.LocalRecord

	if len(localRecordList) > 0 {
		lastLocalRecord = localRecordList[len(localRecordList)-1]
	}

	var diffRemoteRecordList []model.RemoteRecord
	respData, err := FetchRemoteData(gameInfo.GachaUrl, gameInfo.AccessToken, "", poolType)
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
		respData, err = FetchRemoteData(gameInfo.GachaUrl, gameInfo.AccessToken, respData.Next, poolType)
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
		err = SaveLocalRecord(gameInfo.Uid, diffLocalRecordList)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func IncrementalUpdatePoolInfo() (string, error) {
	gameInfo, err := util.GetGameInfo()
	if err != nil {
		return "", errors.WithStack(err)
	}

	err = incrementalUpdatePoolInfo(gameInfo, 1)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(gameInfo, 3)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(gameInfo, 4)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(gameInfo, 5)
	if err != nil {
		return "", errors.WithStack(err)
	}
	err = incrementalUpdatePoolInfo(gameInfo, 8)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return gameInfo.Uid, nil
}
