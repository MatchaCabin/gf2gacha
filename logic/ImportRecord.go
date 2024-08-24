package logic

import (
	"encoding/json"
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"os"
)

func ImportRawJson(uid, rawJsonPath string) error {
	rawJsonBytes, err := os.ReadFile(rawJsonPath)
	if err != nil {
		return errors.WithStack(err)
	}

	rawMap := make(map[int64][]model.RemoteRecord)
	err = json.Unmarshal(rawJsonBytes, &rawMap)
	if err != nil {
		return errors.WithStack(err)
	}

	//解析成功后先做备份
	err = util.BackupDB()
	if err != nil {
		return errors.WithStack(err)
	}

	for poolType, rawRecordList := range rawMap {
		var recordList []model.LocalRecord
		for _, record := range rawRecordList {
			recordList = append(recordList, model.LocalRecord{
				PoolType:       poolType,
				PoolId:         record.PoolId,
				ItemId:         record.ItemId,
				GachaTimestamp: record.GachaTimestamp,
			})
		}

		localRecordList, err := GetLocalRecord(uid, poolType, 0)
		if err != nil {
			return errors.WithStack(err)
		}

		newRecordList := MergeRecord(localRecordList, recordList)
		if len(newRecordList) == 0 {
			return nil
		}

		err = RemoveLocalRecord(uid, poolType)
		if err != nil {
			return errors.WithStack(err)
		}

		err = SaveLocalRecord(uid, newRecordList)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
