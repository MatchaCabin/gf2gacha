package logic

import (
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/pkg/errors"
	"os"
	"strings"
)

func MergeEreRecord(uid, erePath, typ string) error {
	ereFileData, err := os.ReadFile(erePath)
	if err != nil {
		return errors.WithStack(err)
	}

	var ereRecordList []model.LocalRecord

	switch strings.ToLower(typ) {
	case "json":
		ereRecordList, err = ParseEreJsonData(ereFileData)
		if err != nil {
			return errors.WithStack(err)
		}
	case "excel":
		ereRecordList, err = ParseEreExcelData(ereFileData)
		if err != nil {
			return errors.WithStack(err)
		}
	default:
		return errors.Errorf("unknown ere type: %s", typ)
	}

	//解析成功后先做备份
	err = util.BackupDB()
	if err != nil {
		return errors.WithStack(err)
	}

	//需要每个池子分开处理
	poolMap := make(map[int64][]model.LocalRecord)
	for i, record := range ereRecordList {
		poolMap[record.PoolType] = append(poolMap[record.PoolType], ereRecordList[i])
	}

	for poolType, recordList := range poolMap {
		err = mergeEreRecord(uid, poolType, recordList)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func mergeEreRecord(uid string, poolType int64, ereRecordList []model.LocalRecord) error {
	localRecordList, err := GetLocalRecord(uid, poolType, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	newRecordList := MergeRecord(localRecordList, ereRecordList)
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

	return nil
}
